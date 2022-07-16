package model

import (
	"time"
)

// Model defines base gorm model
type Model struct {
	ID uint `gorm:"primary_key" json:"id"`

	DeletedAt *time.Time `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	CreatedAt time.Time  `json:"-"`
}
