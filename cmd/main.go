package main

import (
	application "github.com/CaioAureliano/bank-transaction/internal/config"
	"github.com/CaioAureliano/bank-transaction/internal/config/migration"
)

func init() {
	migration.Migrate()
}

func main() {
	application.Start()
}
