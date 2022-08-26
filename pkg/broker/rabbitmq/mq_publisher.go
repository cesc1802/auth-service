package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/cesc1802/auth-service/pkg/broker"
	"github.com/streadway/amqp"
)

type mqPublisher struct {
	Producer
}

func NewMQPublisher(config MQConfig) *mqPublisher {
	rmq := New(
		&config.Config,
	)

	publisher, err := rmq.NewProducer(config.Exchange, config.Queue, config.PublishingOptions)
	if err != nil {
		panic(err)
	}

	publisher.RegisterSignalHandler()

	publisher.NotifyReturn(func(message amqp.Return) {
		fmt.Println(message)
	})

	return &mqPublisher{
		*publisher,
	}
}

func (p *mqPublisher) Produce(topic string, msg broker.Message) error {
	byteVal, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.Publish(amqp.Publishing{
		Body: byteVal,
	})
}

func (p *mqPublisher) Close() {
	p.Shutdown()
}
