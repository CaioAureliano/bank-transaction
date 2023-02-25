package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_USER     = os.Getenv("DB_USER")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
)

func Connection() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@/%s", DB_USER, DB_PASSWORD, DB_NAME)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return db
}
