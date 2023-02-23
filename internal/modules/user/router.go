package user

import "github.com/gofiber/fiber/v2"

const (
	USER_ENDPOINT = "/user"
)

var h = NewHandler()

func Router(group fiber.Router) {
	group.Post(USER_ENDPOINT, h.CreateUser())
}
