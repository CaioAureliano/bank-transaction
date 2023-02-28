package migration

import (
	"log"

	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

func Migrate() {

	log.Println("attempting to migrate tables")

	dialector, connection := database.DefaultDialector()
	defer connection.Close()

	db := database.Connection(dialector)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Transaction{})
}
