package router

import (
	user "github.com/CaioAureliano/bank-transaction/internal/modules/user/handler"
	"github.com/gofiber/fiber/v2"
)

const (
	DEFAULT_PATH = "/v1"
)

func Router(app *fiber.App) {
	v1 := app.Group(DEFAULT_PATH)

	user.Router(v1)
}
