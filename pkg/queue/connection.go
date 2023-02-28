package queue

import (
	"log"

	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	rabbitmq "github.com/rabbitmq/amqp091-go"
)

func Connection() *rabbitmq.Connection {
	conn, err := rabbitmq.Dial(configuration.Env.QUEUEURI)
	if err != nil {
		log.Panic(err)
	}
	return conn
}

func Close(connection *rabbitmq.Connection) {
	r := recover()
	if r != nil && connection != nil {
		log.Println("closing rabbitmq connection", r)
	}
	connection.Close()
}
