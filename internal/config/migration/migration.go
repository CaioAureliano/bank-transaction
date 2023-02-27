package migration

import (
	"log"

	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

func Migrate() {

	log.Println("attempting to migrate tables")

	db := database.Connection(database.DefaultDialector())

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Transaction{})
}
