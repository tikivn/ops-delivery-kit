package consumer

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/getsentry/raven-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/tag"
)

type BaseMessage struct {
	RequestID   uuid.UUID   `json:"request_id"`
	RequestTime string      `json:"request_time"`
	Source      string      `json:"source"`
	Payload     interface{} `json:"payload"`
}

type Consumer interface {
	MustStart(processor Processor)
}

type Processor interface {
	Topics(context.Context) []string
	Decode(context.Context, []byte) (interface{}, error)
	Process(context.Context, interface{}) error
	String() string
}

func NewConsumer(
	ctx context.Context,
	groupID string,
	brokers []string,
	tracer Tracer,
	whitelistErrors []error,
) Consumer {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Consumer.Return.Errors = true

	if tracer == nil {
		tracer = &noopTracing{}
	}

	return &saramaConsumer{
		ctx:             ctx,
		brokers:         brokers,
		cfg:             config,
		groupID:         groupID,
		tracer:          tracer,
		whitelistErrors: whitelistErrors,
	}
}

type saramaConsumer struct {
	ctx             context.Context
	brokers         []string
	cfg             *sarama.Config
	groupID         string
	tracer          Tracer
	whitelistErrors []error
}

func (c *saramaConsumer) Start(processor Processor) error {
	client, err := sarama.NewClient(c.brokers, c.cfg)
	if err != nil {
		return errors.Wrap(err, "create client fail")
	}
	defer client.Close()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(c.groupID, client)
	if err != nil {
		return errors.Wrap(err, "create group fail")
	}
	defer group.Close()

	// Track errors
	go func() {
		for err := range group.Errors() {
			logrus.Error("error occurs in consumer: ", err)
			if err, ok := err.(*sarama.ConsumerError); ok {
				raven.CaptureError(err.Err, map[string]string{
					"partition": fmt.Sprint(err.Partition),
					"topic":     err.Topic,
				})
			} else {
				raven.CaptureError(err, nil)
			}
		}
	}()

	topics := processor.Topics(c.ctx)

	h := newGroupHandler(NewProcessor(processor, c.tracer), c.tracer, c.whitelistErrors)
	for {
		err := group.Consume(c.ctx, topics, h)
		if err == sarama.ErrClosedClient || err == sarama.ErrClosedConsumerGroup {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (s *saramaConsumer) MustStart(processor Processor) {
	if err := s.Start(processor); err != nil {
		panic(err)
	}
}

type groupHandler struct {
	processor       Processor
	tracer          Tracer
	whitelistErrors []error
}

func newGroupHandler(c Processor, tracer Tracer, whitelistErrors []error) sarama.ConsumerGroupHandler {
	return &groupHandler{
		processor:       c,
		tracer:          tracer,
		whitelistErrors: whitelistErrors,
	}
}

func (c *groupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *groupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c *groupHandler) ConsumeClaim(s sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		err := (func() (err error) {
			ctx, err := tag.New(s.Context(),
				tag.Insert(keyProcessor, c.processor.String()),
				tag.Insert(keyTopic, msg.Topic),
			)
			defer recordMetrics(ctx, msg.Value)(err)

			ctx, done := c.tracer.Transaction(ctx, c.processor.String())
			defer func() {
				done(err)
			}()

			if len(msg.Value) == 0 {
				return nil
			}

			data, err := c.processor.Decode(ctx, msg.Value)
			if err != nil {
				return errors.Wrapf(err,
					"decode message error at offset %d topic %s partition %d, claims: %v", msg.Offset, msg.Topic, msg.Partition, string(msg.Value))
			}

			err = c.processor.Process(ctx, data)
			if err != nil {
				for _, whitelistError := range c.whitelistErrors {
					if strings.Contains(err.Error(), whitelistError.Error()) {
						logrus.Warningf("process message error at offset %d topic %s partition %d, claims: %v", msg.Offset, msg.Topic, msg.Partition, string(msg.Value))
						return nil
					}
				}

				return errors.Wrapf(err,
					"process message error at offset %d topic %s partition %d, claims: %v", msg.Offset, msg.Topic, msg.Partition, string(msg.Value))
			}
			return nil
		})()
		if err != nil {
			return err
		}

		s.MarkMessage(msg, "")
	}

	return nil
}
