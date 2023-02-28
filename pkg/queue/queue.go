package queue

import (
	"context"
	"time"

	rabbitmq "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	conn *rabbitmq.Connection
	ch   *rabbitmq.Channel
}

func New() Queue {
	conn, ch := Channel()
	return Queue{conn, ch}
}

func (b Queue) SendMessage(payload string) error {
	defer Close(b.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return b.ch.PublishWithContext(ctx,
		"",
		"transactions",
		false,
		false,
		rabbitmq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)
}
