package router

import (
	"github.com/gofiber/fiber/v2"
)

const (
	DEFAULT_PATH = "/v1"
)

func Router(app *fiber.App) fiber.Router {
	return app.Group(DEFAULT_PATH)
}
