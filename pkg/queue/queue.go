package queue

import (
	"context"
	"time"

	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
}

func New() Queue {
	return Queue{}
}

func (b Queue) SendMessage(payload string) error {
	conn, channel := Channel()
	defer Close(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return channel.PublishWithContext(ctx,
		"",
		configuration.Env.QUEUENAME,
		false,
		false,
		rabbitmq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)
}
