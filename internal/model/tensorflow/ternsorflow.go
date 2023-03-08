package model

import (
	m "seyes-core/internal/model"
)

// Tensorflows defines the Tensorflows model details
type Tensorflows struct {
	m.Model
	Label        string `json:"label"`
	// ModelFile    []
}
