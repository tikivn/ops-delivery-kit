package publisher

import (
	"context"

	"github.com/micro/go-micro/broker"
	"github.com/tikivn/ops-delivery/pkg/publisher/rabbitmq-plugin" // Modify https://github.com/micro/go-plugins/broker/rabbitmq for default AMQP exchange
)

type RabbitmqPublisher interface {
	Connect() error
	Init() error
	Publish(ctx context.Context, message *broker.Message) error
}

type rabbitmqPublisher struct {
	rabbitmq broker.Broker
	topic    string
}

func NewRabbitmqPublisher(addrs string, exchange string, topic string) RabbitmqPublisher {
	var rbroker broker.Broker

	if exchange == "" { // Use default exchange, default exchange name is empty string
		rbroker = rabbitmq.NewBroker(broker.Addrs(addrs))
	} else {
		rbroker = rabbitmq.NewBroker(broker.Addrs(addrs), rabbitmq.Exchange(exchange))
	}
	return &rabbitmqPublisher{
		rabbitmq: rbroker,
		topic:    topic,
	}
}

func (r *rabbitmqPublisher) Init() error {
	err := r.rabbitmq.Init()
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitmqPublisher) Connect() error {
	err := r.rabbitmq.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitmqPublisher) Publish(ctx context.Context, message *broker.Message) error {
	err := r.rabbitmq.Publish(r.topic, message, rabbitmq.DeliveryMode(2))
	if err != nil {
		return err
	}
	return nil
}
