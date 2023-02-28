package transfer

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/handler"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/repository"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/service"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/worker"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cache *redis.Client) {

	r := repository.New(db, cache)

	s := service.New(r)

	h := handler.New(s)

	go worker.Start(h)
}
