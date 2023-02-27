package config

import (
	"os"

	"github.com/CaioAureliano/bank-transaction/internal/config/router"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/CaioAureliano/bank-transaction/pkg/authentication"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
)

var (
	port = os.Getenv("PORT")
)

func Start() {
	app := api.Setup()
	v1 := router.Router(app)
	db := database.Connection(database.DefaultDialector())

	account.Setup(v1, db)
	app.Use(authentication.JwtMiddleware())
	transaction.Setup(v1, db)

	app.Listen(port)
}
