package broker

import (
	"context"
	"time"

	rabbitmq "github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	conn *rabbitmq.Connection
	ch   *rabbitmq.Channel
	q    rabbitmq.Queue
}

func New() Broker {
	conn, ch, q := createQueue()
	return Broker{conn, ch, q}
}

func (b Broker) Pub(body string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer b.conn.Close()

	return b.ch.PublishWithContext(ctx,
		"",
		b.q.Name,
		false,
		false,
		rabbitmq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}
