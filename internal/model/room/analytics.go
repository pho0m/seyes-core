package model

import (
	m "seyes-core/internal/model"
)

// Analytic defines the analytics model details
type Analytic struct {
	m.Model

	PersonSumaryCont int64   `json:"person_summary_count"`
	ComOnSumaryCount int64   `json:"comon_summary_count"`
	AccurencySummary float64 `json:"accurency_summary"`
}
