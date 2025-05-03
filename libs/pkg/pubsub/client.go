package pubsub

type MessageQueueClient interface {
	Publish(queueName string)
	Subscribe(queueName string)
}
