package queue

import (
	"log"

	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

func Channel() (*rabbitmq.Connection, *rabbitmq.Channel) {

	conn := Connection()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	_, err = ch.QueueDeclare(configuration.Env.QUEUENAME, false, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}

	return conn, ch
}
