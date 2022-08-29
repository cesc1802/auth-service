package broker

import "github.com/streadway/amqp"

type Publisher interface {
	Produce(topic string, msg Message) error
	Close()
}

type Subscriber interface {
	Subscribe(topic string, handler func(delivery amqp.Delivery)) error
	Close()
}

type PubSub interface {
	Publisher
	Subscriber
}
