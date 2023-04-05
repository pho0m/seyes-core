package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	core "seyes-core/internal/core/notifications"
	"seyes-core/internal/helper"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	mo "seyes-core/internal/model/room"
)

// SettingsParams define params for create room
type SettingsParams struct {
	ID                    int64  `json:"id"`
	Active                bool   `json:"active"`
	AiModelData           string `json:"model_data"`
	CronjobTime           string `json:"cronjob_time"`
	LineNotifyAccessToken string `json:"notify_access_token"`
	MqttIp                string `json:"mqtt_ip"`
	MqttUserName          string `json:"mqtt_username"`
	MqttPassword          string `json:"mqtt_password"`
	MqttPort              string `json:"mqtt_port"`
	MqttClientName        string `json:"mqtt_client_name"`
}

// GetSetting get a room by room id
func GetSetting(db *gorm.DB, ps *helper.UrlParams) (map[string]interface{}, error) {
	var setting mo.Setting

	if err := db.First(&setting).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                  setting.ID,
		"active":              setting.Active,
		"model_data":          setting.AiModelData,
		"cronjob_time":        setting.CronjobTime,
		"notify_access_token": setting.LineNotifyAccessToken,
		"update_at":           setting.UpdatedAt,
		"mqtt_ip":             setting.MqttIp,
		"mqtt_username":       setting.MqttUserName,
		"mqtt_password":       setting.MqttPassword,
		"mqtt_port":           setting.MqttPort,
		"mqtt_client_name":    setting.MqttClientName,
	}

	return res, nil
}

// CreateSettings create a room
func CreateSettings(db *gorm.DB, ps *SettingsParams) (map[string]interface{}, error) {
	setting := &mo.Setting{
		Active:                ps.Active,
		AiModelData:           ps.AiModelData,
		CronjobTime:           ps.CronjobTime,
		LineNotifyAccessToken: ps.LineNotifyAccessToken,
		MqttIp:                ps.MqttIp,
		MqttUserName:          ps.MqttUserName,
		MqttPassword:          ps.MqttPassword,
		MqttPort:              ps.MqttPort,
		MqttClientName:        ps.MqttClientName,
	}

	if err := db.Create(&setting).Error; err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                  setting.ID,
		"active":              setting.Active,
		"model_data":          setting.AiModelData,
		"cronjob_time":        setting.CronjobTime,
		"notify_access_token": setting.LineNotifyAccessToken,
		"update_at":           setting.UpdatedAt,
		"mqtt_ip":             setting.MqttIp,
		"mqtt_username":       setting.MqttUserName,
		"mqtt_password":       setting.MqttPassword,
		"mqtt_port":           setting.MqttPort,
		"mqtt_client_name":    setting.MqttClientName,
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

	cr := cron.New()
	cr.Start()

	spew.Dump(cr)
	spew.Dump("test update setting")

	setting.Active = ps.Active
	setting.AiModelData = ps.AiModelData
	setting.CronjobTime = ps.CronjobTime
	setting.LineNotifyAccessToken = ps.LineNotifyAccessToken
	setting.MqttIp = ps.MqttIp
	setting.MqttUserName = ps.MqttUserName
	setting.MqttPassword = ps.MqttPassword
	setting.MqttPort = ps.MqttPort
	setting.MqttClientName = ps.MqttClientName

	if err := db.Save(&setting).Error; err != nil {
		return nil, err
	}

	cronAutomateDetection(db, setting.CronjobTime, cr) // current
	cr.Start()

	res := map[string]interface{}{
		"id":                  setting.ID,
		"active":              setting.Active,
		"model_data":          setting.AiModelData,
		"cronjob_time":        setting.CronjobTime,
		"notify_access_token": setting.LineNotifyAccessToken,
		"update_at":           setting.UpdatedAt,
		"mqtt_ip":             setting.MqttIp,
		"mqtt_username":       setting.MqttUserName,
		"mqtt_password":       setting.MqttPassword,
		"mqtt_port":           setting.MqttPort,
		"mqtt_client_name":    setting.MqttClientName,
	}

	return res, nil
}

func cronAutomateDetection(db *gorm.DB, cronTime string, cr *cron.Cron) {

	cr.AddFunc("@every "+cronTime+"m", func() {
		spew.Dump(cronTime)
		spew.Dump("Start cron !")
		// ctx := r.Context().Value("user_info").(*auth.UserInfo)
		var urlSeyesCam = os.Getenv("SEYES_CAM_URL") + "/image"

		resRoom, err := GetAllRoom(db, &RoomFilter{
			ID:      0,
			Page:    1,
			OrderBy: "",
			SortBy:  "",
		})

		if err != nil {
			return
		}

		rr, _ := json.Marshal(resRoom)
		var rooms []RoomParams
		json.Unmarshal(rr, &rooms)

		for _, s := range rooms {

			resFromSCAM, err := http.Get(urlSeyesCam + "/" + s.UuidCam + "/channel" + "/0") //+ "/" + ps.Uuid + "/channel" + ps.Channel
			if err != nil {
				return
			}

			responseData, err := ioutil.ReadAll(resFromSCAM.Body)
			if err != nil {
				return
			}
			var dp DetectionParams
			json.Unmarshal(responseData, &dp)

			var jsonStr = []byte(`{"image":` + `"` + dp.ImageData + `"` + `}`)
			req, err := http.NewRequest("POST", os.Getenv("SEYES_DETECT_URL")+"/detect", bytes.NewBuffer(jsonStr))
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &dp)

			personCount, _ := strconv.Atoi(dp.PersonCount)
			comCount, _ := strconv.Atoi(dp.ConOnCount)
			status := "undetected"

			if dp.Accurency != "0%" {
				status = "detected"
			}

			resReports, err := CreateReport(db, &ReportsParams{
				PersonCont: int64(personCount),
				ComOnCount: int64(comCount),
				Status:     status,
				Image:      dp.ImageData,
				Accurency:  dp.Accurency,
				RoomLabel:  s.Label,
				ReportTime: dp.Time,
				ReportDate: dp.Date,
			})

			if err != nil {
				return
			}

			rreport, _ := json.Marshal(resReports)
			var report ReportsParams
			json.Unmarshal(rreport, &report)

			notifyData := core.NotifyParamV2{
				Uuid:      s.UuidCam,
				Image:     report.Image,
				ID:        report.ID,
				Person:    strconv.Itoa(int(report.PersonCont)),
				ComOn:     strconv.Itoa(int(report.ComOnCount)),
				UploadAt:  report.ReportDate,
				Time:      report.ReportTime,
				Accurency: report.Accurency,
			}

			if notifyData.Accurency != "0%" {
				err = core.SendToLineNotifyV2(&notifyData)

				if err != nil {
					return
				}
			}

		}

	},
	)
}
