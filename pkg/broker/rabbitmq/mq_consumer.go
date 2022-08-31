package rabbitmq

import "github.com/streadway/amqp"

type mqConsumer struct {
	Consumer
}

func NewMQConsumer(config MQConfig) *mqConsumer {
	rmq := New(
		&config.Config,
	)

	consumer, err := rmq.NewConsumer(config.Exchange, config.Queue, config.BindingOptions, config.ConsumerOptions)
	if err != nil {
		panic(err)
	}
	// defer consumer.Shutdown()
	err = consumer.QOS(3)
	if err != nil {
		panic(err)
	}

	consumer.RegisterSignalHandler()
	return &mqConsumer{
		*consumer,
	}
}

func (c *mqConsumer) Subscribe(handler func(delivery amqp.Delivery)) error {
	return c.Consume(handler)
}

func (c *mqConsumer) Close() {
	c.Consumer.Shutdown()
}
