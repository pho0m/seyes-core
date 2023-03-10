package model

import (
	m "seyes-core/internal/model"
)

// Schedule defines the schedule model details
type Schedule struct {
	m.Model

	StartTime string `json:"start_time"`
	DueTime   string `json:"due_time"`
	Period    string `json:"period"`
	Day       string `json:"day"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
	Label     string `json:"label"`
	Prefix    string `json:"prefix"`
	Active    bool   `json:"active"`
}
