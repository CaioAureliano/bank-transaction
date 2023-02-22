package user

import "github.com/gofiber/fiber/v2"

const (
	endpoint = "/user"
)

var h = NewHandler()

func Router(group fiber.Router) {
	group.Post(endpoint, h.CreateUser())
}
