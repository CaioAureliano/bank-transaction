package router

import (
	account "github.com/CaioAureliano/bank-transaction/internal/modules/account/handler"
	"github.com/gofiber/fiber/v2"
)

const (
	DEFAULT_PATH = "/v1"
)

func Router(app *fiber.App) {
	v1 := app.Group(DEFAULT_PATH)
	account.Router(v1)
}
