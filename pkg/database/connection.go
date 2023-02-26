package database

import (
	"log"

	"gorm.io/gorm"
)

func Connection(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return db
}
