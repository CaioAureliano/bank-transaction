package queue

import (
	"log"
	"os"

	rabbitmq "github.com/rabbitmq/amqp091-go"
)

var (
	RABBITMQ_QUEUE_NAME = os.Getenv("RABBITMQ_QUEUE_NAME")
)

func Channel() (*rabbitmq.Connection, *rabbitmq.Channel) {

	conn := Connection()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	_, err = ch.QueueDeclare("transactions", false, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}

	return conn, ch
}
