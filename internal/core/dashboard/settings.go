package core

import (
	"seyes-core/internal/helper"

	"gorm.io/gorm"

	mo "seyes-core/internal/model/room"
)

// SettingsParams define params for create room
type SettingsParams struct {
	ID                    int64  `json:"id"`
	AiModelData           string `json:"model_data"`
	CronjobTime           string `json:"cronjob_time"`
	LineNotifyAccessToken string `json:"notify_access_token"`
	MqttIp                string `json:"mqtt_ip"`
	MqttUserName          string `json:"mqtt_username"`
	MqttPassword          string `json:"mqtt_password"`
	MqttPort              string `json:"mqtt_port"`
}

// GetSetting get a room by room id
func GetSetting(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	var setting mo.Setting

	if err := db.First(&setting).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                  setting.ID,
		"model_data":          setting.AiModelData,
		"cronjob_time":        setting.CronjobTime,
		"notify_access_token": setting.LineNotifyAccessToken,
		"update_at":           setting.UpdatedAt,
		"mqtt_ip":             setting.MqttIp,
		"mqtt_username":       setting.MqttUserName,
		"mqtt_password":       setting.MqttPassword,
		"mqtt_port":           setting.MqttPort,
	}

	return res, nil
}

// CreateSettings create a room
func CreateSettings(db *gorm.DB, ps *SettingsParams) (map[string]interface{}, error) {
	setting := &mo.Setting{
		AiModelData:           ps.AiModelData,
		CronjobTime:           ps.CronjobTime,
		LineNotifyAccessToken: ps.LineNotifyAccessToken,
		MqttIp:                ps.MqttIp,
		MqttUserName:          ps.MqttUserName,
		MqttPassword:          ps.MqttPassword,
		MqttPort:              ps.MqttPort,
	}

	if err := db.Create(&setting).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                  setting.ID,
		"model_data":          setting.AiModelData,
		"cronjob_time":        setting.CronjobTime,
		"notify_access_token": setting.LineNotifyAccessToken,
		"update_at":           setting.UpdatedAt,
		"mqtt_ip":             setting.MqttIp,
		"mqtt_username":       setting.MqttUserName,
		"mqtt_password":       setting.MqttPassword,
		"mqtt_port":           setting.MqttPort,
	}

	return res, nil
}

// UpdatedSettings update a room
func UpdatedSettings(db *gorm.DB, ps *SettingsParams) (map[string]interface{}, error) {
	var setting mo.Setting

	if err := db.Where("id = ?", ps.ID).
		First(&setting).Error; err != nil {
		return nil, err
	}

	setting.AiModelData = ps.AiModelData
	setting.CronjobTime = ps.CronjobTime
	setting.LineNotifyAccessToken = ps.LineNotifyAccessToken
	setting.MqttIp = ps.MqttIp
	setting.MqttUserName = ps.MqttUserName
	setting.MqttPassword = ps.MqttPassword
	setting.MqttPort = ps.MqttPort

	if err := db.Save(&setting).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                  setting.ID,
		"model_data":          setting.AiModelData,
		"cronjob_time":        setting.CronjobTime,
		"notify_access_token": setting.LineNotifyAccessToken,
		"update_at":           setting.UpdatedAt,
		"mqtt_ip":             setting.MqttIp,
		"mqtt_username":       setting.MqttUserName,
		"mqtt_password":       setting.MqttPassword,
		"mqtt_port":           setting.MqttPort,
	}

	return res, nil
}
