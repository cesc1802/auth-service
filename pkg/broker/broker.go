package broker

type Publisher interface {
}

type Subscriber interface {
}

type PubSub interface {
	Publisher
	Subscriber
}
