package rabbitmq

type MQConfig struct {
	Config            Config
	Exchange          Exchange
	Queue             Queue
	BindingOptions    BindingOptions
	ConsumerOptions   ConsumerOptions
	PublishingOptions PublishingOptions
}
