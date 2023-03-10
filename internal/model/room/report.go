package model

import (
	m "seyes-core/internal/model"
)

// Report defines the log notification model details
type Report struct {
	m.Model

	PersonCont int64   `json:"person_count"`
	ComOnCount int64   `json:"com_on_count"`
	Accurency  float64 `json:"accurency"`
	RoomLabel  string  `json:"room_label"`
	ReportTime string  `json:"report_time"`
	ReportDate string  `json:"report_date"`
	Photo      string  `json:"photo"`
}
