package service

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB creates a new DB service
func NewDB() (*gorm.DB, error) {
	c := &gorm.Config{}
	dsn := databaseDSN()

	if os.Getenv("APP_ENV") == "prod" {
		c.Logger = logger.Default.LogMode(logger.Silent)
	}

	return gorm.Open(postgres.Open(dsn), c)
}

// databaseDSN dsn value
func databaseDSN() string {
	return os.Getenv("DB_DSN")
}
