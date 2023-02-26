package account

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/handler"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/repository"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/service"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, group fiber.Router, db *gorm.DB) {

	r := repository.New(db)

	v := validator.New(r)

	s := service.New(r, v)

	h := handler.New(s)

	handler.Router(group, h)
}
