package handler

import "github.com/gofiber/fiber/v2"

const (
	USER_ENDPOINT = "/user"
)

func Router(group fiber.Router) {
	group.Post(USER_ENDPOINT, CreateUser)
}
