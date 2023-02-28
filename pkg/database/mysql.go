package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DefaultDialector() (gorm.Dialector, *sql.DB) {
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", configuration.Env.DBUSER, configuration.Env.DBPASSWORD, configuration.Env.DBNAME)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	return mysql.New(mysql.Config{Conn: conn}), conn
}
