package database

import (
	"os"

	"gorm.io/gorm/logger"
)

// Get logger interface by ENV variable, silent to production("PROD")
func getLogger() logger.Interface {
	sqlLogger := logger.Default.LogMode(logger.Error)
	if value, ok := os.LookupEnv("ENV"); ok && value != "" {
		sqlLogger = logger.Default.LogMode(logger.Silent)
	}
	return sqlLogger
}
