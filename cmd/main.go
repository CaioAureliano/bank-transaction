package main

import (
	_ "github.com/CaioAureliano/bank-transaction/docs"
	application "github.com/CaioAureliano/bank-transaction/internal/config"
	"github.com/CaioAureliano/bank-transaction/internal/config/migration"
)

func init() {
	migration.Migrate()
}

// @title Bank Transaction
// @version 1.0
// @description A simple Restful API to Bank Transaction

// @BasePath	/v1/

// @securityDefinitions.apikey	JwtToken
// @in							header
// @name						Authorization
// @description					use with "Bearer " prefix. e.g: "Authorization: Bearer {token}"
func main() {
	application.Start()
}
