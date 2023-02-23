package server

import (
	"log"
	"os"

	"github.com/CaioAureliano/bank-transaction/internal/config/router"
	"github.com/CaioAureliano/bank-transaction/internal/shared/application"
)

var (
	port = os.Getenv("PORT")
)

func Run() {
	app := application.Setup()
	router.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
