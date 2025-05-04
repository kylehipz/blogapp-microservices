package pubsub

import "context"

type Message struct {
	Event   string
	Payload any
}

type PubSubClient interface {
	Publish(ctx context.Context, event string, payload any) error
	Subscribe(events []string) (<-chan *Message, error)
}
