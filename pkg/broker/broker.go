package broker

import "github.com/streadway/amqp"

type Publisher interface {
	Produce(msg Message) error
	Close()
}

type Subscriber interface {
	Subscribe(handler func(delivery amqp.Delivery)) error
	Close()
}

type PubSub interface {
	Publisher
	Subscriber
}
