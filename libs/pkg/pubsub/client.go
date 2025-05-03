package pubsub

type Message struct {
	Event   string
	Payload any
}

type MessageQueueClient interface {
	Publish(queueName string)
	Subscribe(queueName string)
}
