package model

import (
	m "mns-core/internal/model"
)

// Shop defines the shop details
type Shop struct {
	m.Model
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}
