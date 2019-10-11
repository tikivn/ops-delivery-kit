package publisher

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type MarshalFunc func(interface{}) ([]byte, error)

type KafkaPublisher struct {
	p       sarama.SyncProducer
	topic   string
	marshal MarshalFunc
}

func NewPublisher(brokers []string, topic string) (*KafkaPublisher, error) {
	p, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		return nil, err
	}
	return &KafkaPublisher{
		p:       p,
		topic:   topic,
		marshal: json.Marshal,
	}, nil
}

func (p *KafkaPublisher) Publish(ctx context.Context, key []byte, value interface{}) error {
	bytes, err := p.marshal(value)
	if err != nil {
		return errors.Wrapf(err, "Publisher: marshal error (%s)", key)
	}
	_, _, err = p.p.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(bytes),
	})
	if err != nil {
		return errors.Wrapf(err, "Publisher: Publish message error (%s)", key)
	}
	return nil
}

func (p *KafkaPublisher) Close() error {
	return p.p.Close()
}
