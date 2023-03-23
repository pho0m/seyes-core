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
	Image      string  `json:"image"`

	Status string `json:"status"`
	Lamp1  string `json:"lamp_1_status"`
	Lamp2  string `json:"lamp_2_status"`
	Lamp3  string `json:"lamp_3_status"`
	Lamp4  string `json:"lamp_4_status"`
	Lamp5  string `json:"lamp_5_status"`
	Lamp6  string `json:"lamp_6_status"`
	Door   string `json:"door_status"`
	Air    string `json:"air_status"`
}
