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

func mysqlDialector() gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@/%s", DB_USER, DB_NAME, DB_PASSWORD)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	return mysql.New(mysql.Config{Conn: conn})
}
