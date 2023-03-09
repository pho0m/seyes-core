package model

import (
	m "seyes-core/internal/model"
)

// Room defines the log notification model details
type Room struct {
	m.Model
	Label  string `json:"label"`
	CamURL string `json:"cam_url"`
	Status string `json:"status"`
	Active string `json:"active"`
}
