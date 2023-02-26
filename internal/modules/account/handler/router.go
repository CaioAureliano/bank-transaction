package handler

import "github.com/gofiber/fiber/v2"

type handlers interface {
	CreateUser(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
}

const (
	accountEndpoint = "/accounts"
)

func Router(group fiber.Router, h handlers) {
	group.Post(accountEndpoint, h.CreateUser)
	group.Post(accountEndpoint+"/auth", h.Authenticate)
}
