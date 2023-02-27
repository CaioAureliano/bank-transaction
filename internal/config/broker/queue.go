package broker

import (
	"log"
	"os"

	"github.com/CaioAureliano/bank-transaction/pkg/queue"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

var (
	queueName = os.Getenv("RABBITMQ_QUEUE_NAME")
)

func createQueue() (*rabbitmq.Connection, *rabbitmq.Channel, rabbitmq.Queue) {

	conn := queue.Connection()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panic(err)
	}

	return conn, ch, q
}
