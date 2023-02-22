package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	port = os.Getenv("PORT")
)

func Run() {
	app := fiber.New()
	app.Listen(port)
}
