package model

import (
	m "seyes-core/internal/model"
)

// Log defines the log notification model details
type Log struct {
	m.Model
	Person int64 `json:"person"`
	ComOn  int64 `json:"com_on"`
}
