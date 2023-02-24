package main

import (
	"os"

	"github.com/CaioAureliano/bank-transaction/internal/config/router"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	app := api.Setup()
	router.Router(app)
	app.Listen(port)
}
