package queue

import (
	"log"

	rabbitmq "github.com/rabbitmq/amqp091-go"
)

func Connection() *rabbitmq.Connection {
	conn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672/")
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
