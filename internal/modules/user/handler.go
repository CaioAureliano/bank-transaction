package user

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateUser() fiber.Handler
}

type handler struct{}

func NewHandler() Handler {
	return handler{}
}

func (h handler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("created")
	}
}
