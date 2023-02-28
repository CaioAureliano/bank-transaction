package transaction

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/handler"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/repository"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/service"
	"github.com/CaioAureliano/bank-transaction/pkg/queue"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(group fiber.Router, db *gorm.DB, cache *redis.Client) {

	r := repository.New(db, queue.New(), cache)

	s := service.New(r)

	h := handler.New(s)

	handler.Router(group, h)
}
