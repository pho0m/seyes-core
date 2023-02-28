package model

import (
	m "seyes-core/internal/model"
)

// Schedule defines the schedule model details
type Schedule struct {
	m.Model
	Label        string `json:"label"`
	Prefix       string `json:"prefix"`
	LecturerName string `json:"lecturer_name"`
	StartClass   string `json:"start_class"`
	DueClass     string `json:"due_class"`
	Subject      string `json:"subject"`
}
