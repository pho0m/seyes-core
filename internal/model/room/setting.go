package model

import (
	m "seyes-core/internal/model"
)

// Setting defines the schedule model details
type Setting struct {
	m.Model

	ModelData             string `json:"model_data"`
	CronjobTime           string `json:"cronjob_time"`
	LineNotifyAccessToken string `json:"notify_access_token"`
}
