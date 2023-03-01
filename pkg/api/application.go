package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(idempotency.New())
	return app
}
