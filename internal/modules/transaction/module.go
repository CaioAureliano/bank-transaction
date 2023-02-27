package transaction

import (
	"github.com/CaioAureliano/bank-transaction/internal/config/broker"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/handler"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/repository"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(group fiber.Router, db *gorm.DB) {

	r := repository.New(db, broker.New())

	s := service.New(r)

	h := handler.New(s)

	handler.Router(group, h)
}
