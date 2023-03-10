package model

import (
	m "seyes-core/internal/model"
)

// Setting defines the schedule model details
type Setting struct {
	m.Model

	AiModelData           string `json:"model_data"`
	CronjobTime           string `json:"cronjob_time"`
	LineNotifyAccessToken string `json:"notify_access_token"`

	MqttIp       string `json:"mqtt_ip"`
	MqttUserName string `json:"mqtt_username"`
	MqttPassword string `json:"mqtt_password"`
	MqttPort     string `json:"mqtt_port"`
}
