package model

import (
	m "seyes-core/internal/model"
)

// User defines the user model details
type User struct {
	m.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
	Password  string `json:"password"` //FIXME
	Email     string `json:"email" gorm:"uniqueIndex:idx_user"`
}
