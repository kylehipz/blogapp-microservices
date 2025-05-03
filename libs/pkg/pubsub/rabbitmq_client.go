package pubsub

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
	queue    string
	service  string
}

func NewRabbitMQClient(
	conn *amqp.Connection,
	ch *amqp.Channel,
	exchange string,
	queue string,
	service string,
) *RabbitMQClient {
	err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	return &RabbitMQClient{
		conn:     conn,
		ch:       ch,
		exchange: exchange,
		queue:    queue,
		service:  service,
	}
}

func (r *RabbitMQClient) CleanUp() {
	r.conn.Close()
	r.ch.Close()
}

func (r *RabbitMQClient) Publish(ctx context.Context, event string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = r.ch.PublishWithContext(ctx, r.exchange, event, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQClient) Subscribe(events []string) (<-chan *Message, error) {
	q, err := r.ch.QueueDeclare(r.queue, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		err = r.ch.QueueBind(q.Name, event, r.exchange, false, nil)
		if err != nil {
			return nil, err
		}
	}

	msgs, err := r.ch.Consume(q.Name, r.service, true, false, false, false, nil)

	transformed := make(chan *Message)

	go func() {
		for msg := range msgs {
			var payload any

			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				continue
			}

			transformed <- &Message{
				Event:   msg.RoutingKey,
				Payload: payload,
			}
		}
		close(transformed)
	}()

	return transformed, err
}
