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

	MqttTopicLamp1 string `json:"mqtt_topic_lamp_1"`
	MqttTopicLamp2 string `json:"mqtt_topic_lamp_2"`
	MqttTopicLamp3 string `json:"mqtt_topic_lamp_3"`
	MqttTopicLamp4 string `json:"mqtt_topic_lamp_4"`
	MqttTopicLamp5 string `json:"mqtt_topic_lamp_5"`
	MqttTopicLamp6 string `json:"mqtt_topic_lamp_6"`
	MqttTopicDoor  string `json:"mqtt_topic_door"`
	MqttTopicAir   string `json:"mqtt_topic_air"`
}
