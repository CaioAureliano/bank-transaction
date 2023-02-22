package router

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/user"
	"github.com/gofiber/fiber/v2"
)

const (
	DEFAULT_PATH = "/v1"
)

func Setup(app *fiber.App) {
	v1 := app.Group(DEFAULT_PATH)

	user.Router(v1)
}
