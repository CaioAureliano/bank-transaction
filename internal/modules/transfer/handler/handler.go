package handler

import (
	"encoding/json"
	"log"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain"
)

type service interface {
	Transfer(*domain.TransactionMessage) error
}

type Handler struct {
	s service
}

func New(s service) Handler {
	return Handler{s}
}

func (h Handler) MessageHandler(msg []byte) {

	message := new(domain.TransactionMessage)

	json.Unmarshal(msg, &message)

	log.Printf("received transaction message: %s", msg)

	if err := h.s.Transfer(message); err != nil {
		log.Printf("failed to process message with error: %s", err.Error())
	}
}
