package worker

import (
	"log"

	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/CaioAureliano/bank-transaction/pkg/queue"
)

type handlers interface {
	MessageHandler(msg []byte)
}

func Start(h handlers) {
	conn, channel := queue.Channel()
	defer queue.Close(conn)

	msgs, err := channel.Consume(configuration.Env.QUEUENAME, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("failed to consume with error: %s", err.Error())
	}

	for msg := range msgs {
		h.MessageHandler(msg.Body)
		msg.Ack(false)
	}
}
