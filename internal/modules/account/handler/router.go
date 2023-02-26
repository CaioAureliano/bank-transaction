package handler

import "github.com/gofiber/fiber/v2"

type handlers interface {
	CreateUser(c *fiber.Ctx) error
}

const (
	USER_ENDPOINT = "/user"
)

func Router(group fiber.Router, h handlers) {
	group.Post(USER_ENDPOINT, h.CreateUser)
}
