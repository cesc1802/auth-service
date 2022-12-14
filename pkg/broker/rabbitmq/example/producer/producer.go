package main

import (
	"fmt"
	"github.com/cesc1802/auth-service/pkg/broker/rabbitmq"
	"github.com/cesc1802/auth-service/pkg/logger"
	"strconv"

	"github.com/streadway/amqp"
)

func main() {
	var l = logger.Init(
		logger.WithLogDir("logs/"),
		logger.WithDebug(true),
		logger.WithConsole(true),
	)
	defer l.Sync()

	rmq := rabbitmq.New(
		&rabbitmq.Config{
			Host:     "localhost",
			Port:     5672,
			Username: "guest",
			Password: "guest",
			Vhost:    "/",
		},
	)

	exchange := rabbitmq.Exchange{
		Name: "EXCHANGE_NAME",
	}

	queue := rabbitmq.Queue{}
	publishingOptions := rabbitmq.PublishingOptions{
		Tag:        "ProducerTagHede",
		RoutingKey: "naber",
	}

	publisher, err := rmq.NewProducer(exchange, queue, publishingOptions)
	if err != nil {
		panic(err)
	}
	defer publisher.Shutdown()
	publisher.RegisterSignalHandler()

	//// may be we should autoconvert to byte array?
	//msg := amqp.Publishing{
	//	Body: []byte("2"),
	//}

	publisher.NotifyReturn(func(message amqp.Return) {
		fmt.Println(message)
	})

	for i := 0; i < 10; i++ {
		err = publisher.Publish(amqp.Publishing{
			Body: []byte(strconv.Itoa(i)),
		})
		if err != nil {
			fmt.Println(err, i)
		}
	}
}
