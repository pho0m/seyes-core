package model

import (
	"database/sql"
	m "seyes-core/internal/model"
)

// Room defines the log notification model details
type Room struct {
	m.Model

	Label      string        `json:"label"`
	CamURL     string        `json:"cam_url"`
	Status     string        `json:"status"`
	Active     bool          `json:"active"`
	ScheduleID sql.NullInt64 `json:"schedule_id"`
	Schedule   Schedule
}
