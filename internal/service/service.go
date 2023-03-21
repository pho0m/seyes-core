package service

import (
	"errors"

	"gorm.io/gorm"

	modelRoom "seyes-core/internal/model/room"
	modelUser "seyes-core/internal/model/user"
)

var tableSets = []interface{}{
	modelUser.User{},
	modelRoom.Report{},
	modelRoom.Room{},
	modelRoom.Schedule{},
	modelRoom.Setting{},
}

// Container defines a service container
type Container struct {
	DB   *gorm.DB
	Auth interface{}
}

// NewContainer creates a new service container
func NewContainer() (*Container, error) {
	db, err := NewDB()

	if err != nil {
		return nil, errors.New("failed to initialize database service: " + err.Error())
	}

	return &Container{
		DB:   db,
		Auth: nil,
	}, nil
}

// DoMigration create migration database
func DoMigration(db *gorm.DB) error {
	db.AutoMigrate(tableSets...)

	// if err := CreateDefaultRole(db); err != nil {
	// 	sentry.CaptureException(err)
	// 	panic("cannot initialize Role: " + err.Error())
	// }

	// if err := CreateDefaultSuperAdmin(db); err != nil {
	// 	sentry.CaptureException(err)
	// 	panic("cannot initialize Super admin: " + err.Error())
	// }
	return nil
}
