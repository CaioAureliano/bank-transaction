package database

import (
	"time"

	"gorm.io/gorm"
)

func Connection(dialector gorm.Dialector) *gorm.DB {
	for i := 1; i <= 3; i++ {
		db, err := connect(dialector)
		if err == nil && db != nil {
			return db
		}

		time.Sleep(time.Second * 15)
	}

	panic("failed to initialize database")
}

func connect(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: getLogger(),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
