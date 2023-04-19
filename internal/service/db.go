package service

import (
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	mu "seyes-core/internal/model/user"
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

// CreateDefaultSuperAdmin creating default super admin
func CreateDefaultSuperAdmin(db *gorm.DB) error {
	var user mu.User
	email := os.Getenv("DEFAULT_EMAIL")
	pass := os.Getenv("DEFAULT_PASS")
	bcPass, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		u := mu.User{
			FirstName: "Super",
			LastName: "admin",
			Active: true,
			Tel: "-",
			Email:    email,
			Password: string(bcPass),
		}

		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}

	return nil
}
