package database

import (
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection(dialector gorm.Dialector) *gorm.DB {

	log.Println("attempting to connect to database")

	for i := 1; i <= 3; i++ {
		db, err := connect(dialector)
		if err == nil && db != nil {
			return db
		}

		log.Printf("%d tries to connect", i)

		time.Sleep(time.Second * 10)
	}

	panic("failed to initialize database")
}

func connect(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
