package handler

import "github.com/gofiber/fiber/v2"

type handlers interface {
	CreateTransaction(c *fiber.Ctx) error
}

const (
	transactionEndpoint = "/transactions"
)

func Router(group fiber.Router, h handlers) {
	group.Post(transactionEndpoint, h.CreateTransaction)
}
