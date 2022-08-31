package broker

import (
	"context"

	"github.com/streadway/amqp"
)

type Publisher interface {
	Produce(context.Context, Message) error
	Close()
}

type Subscriber interface {
	Subscribe(context.Context, func(delivery amqp.Delivery)) error
	Close()
}

type PubSub interface {
	Publisher
	Subscriber
}
