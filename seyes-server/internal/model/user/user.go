package model

import (
	m "seyes-core/internal/model"
)

// User defines the user model details
type User struct {
	m.Model
	Name     string `json:"name"`
	Prefix   string `json:"prefix"`
	Password string `json:"password"`
	Phone    string `json:"phone" `
	Email    string `json:"email" gorm:"uniqueIndex:idx_user"`
	Pincode  string `json:"pincode"`
}
