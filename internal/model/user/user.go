package model

import (
	m "seyes-core/internal/model"
)

// User defines the user model details
type User struct {
	m.Model
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}
