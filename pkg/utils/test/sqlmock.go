package test

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func DialectorMock(conn *sql.DB) gorm.Dialector {
	return mysql.New(mysql.Config{
		Conn:                      conn,
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})
}
