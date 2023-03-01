package config

import (
	"github.com/CaioAureliano/bank-transaction/internal/config/router"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/CaioAureliano/bank-transaction/pkg/authentication"
	"github.com/CaioAureliano/bank-transaction/pkg/cache"
	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
	swagger "github.com/arsmn/fiber-swagger/v2"
)

func Start() {
	app := api.Setup()
	v1 := router.Router(app)

	dialector, _ := database.DefaultDialector()
	db := database.Connection(dialector)
	redis := cache.Connection()

	v1.Get("/swagger/*", swagger.HandlerDefault)

	go transfer.Setup(db, redis)
	account.Setup(v1, db)
	app.Use(authentication.JwtMiddleware())
	transaction.Setup(v1, db, redis)

	app.Listen(":" + configuration.Env.PORT)
}
