package server

import (
	"log"
	"os"

	"github.com/CaioAureliano/bank-transaction/internal/config/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	port = os.Getenv("PORT")
)

func Run() {
	app := fiber.New()
	app.Use(cors.New())

	router.Setup(app)

	log.Fatal(app.Listen(port))
}
