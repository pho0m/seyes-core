package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	core "seyes-core/internal/core/dashboard"
	noti "seyes-core/internal/core/notifications"
	model "seyes-core/internal/model/room"

	"seyes-core/internal/helper"
	"seyes-core/internal/service"
	"seyes-core/internal/web"
	"strconv"

	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	if err := loadEnv(); err != nil {
		panic(err)
	}

	appPort := os.Getenv("APP_PORT")

	sc, err := service.NewContainer()

	if err != nil {
		panic(err)
	}

	if err := sanityChecks(); err != nil {
		panic(err)
	}

	s := web.NewServer(sc, appPort)

	if err := service.DoMigration(sc.DB); err != nil {
		panic("cannot initialize Database: " + err.Error())
	}
	initSetting(sc.DB)

	logrus.Info("Starting seyes http server...")
	logrus.Info("Listening in port:" + appPort)

	// cronAutomateDetection(sc.DB)

	s.Start(sc)
}

func cronAutomateDetection(db *gorm.DB) {
	c := cron.New()
	sdata, err := core.GetSetting(db, &helper.UrlParams{ID: 1})

	if err != nil {
		panic(err)
	}
	sm, _ := json.Marshal(sdata)
	var setting core.SettingsParams
	json.Unmarshal(sm, &setting)

	cornTime := setting.CronjobTime

	c.AddFunc("@every "+"30"+"s", func() {

		logrus.Info("[Job] Now Every " + "10" + " job")
		logrus.Info("Room 702 Automate !")

		const urlSeyesCam = "http://202.44.35.76:9093/image"
		resRoom, err := core.GetRoom(db, &helper.UrlParams{
			ID: 1,
		})

		if err != nil {
			return
		}

		rr, _ := json.Marshal(resRoom)
		var room core.RoomParams
		json.Unmarshal(rr, &room)

		resFromSCAM, err := http.Get(urlSeyesCam + "/" + room.UuidCam + "/channel" + "/0") //+ "/" + ps.Uuid + "/channel" + ps.Channel
		if err != nil {
			return
		}

		responseData, err := ioutil.ReadAll(resFromSCAM.Body)
		if err != nil {
			return
		}
		var dp core.DetectionParams
		json.Unmarshal(responseData, &dp)

		var jsonStr = []byte(`{"image":` + `"` + dp.ImageData + `"` + `}`)
		req, err := http.NewRequest("POST", "http://202.44.35.76:9094/detect", bytes.NewBuffer(jsonStr))
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

		resReports, err := core.CreateReport(db, &core.ReportsParams{
			PersonCont: int64(personCount),
			ComOnCount: int64(comCount),
			Status:     "detected",
			Image:      dp.ImageData,
			Accurency:  dp.Accurency,
			RoomLabel:  room.Label,
			ReportTime: dp.Time,
			ReportDate: dp.Date,
		})

		if err != nil {
			return
		}

		rreport, _ := json.Marshal(resReports)
		var report core.ReportsParams
		json.Unmarshal(rreport, &report)

		notifyData := noti.NotifyParamV2{
			Image:     report.Image,
			ID:        report.ID,
			Person:    strconv.Itoa(int(report.PersonCont)),
			ComOn:     strconv.Itoa(int(report.ComOnCount)),
			UploadAt:  report.ReportDate,
			Time:      report.ReportTime,
			Accurency: report.Accurency,
		}

		err = noti.SendToLineNotifyV2(&notifyData)

		if err != nil {
			return
		}

		logrus.Info("[Job] Now Every " + "15" + " job")
		logrus.Info("Room 709 Automate !")

		resRoom2, err := core.GetRoom(db, &helper.UrlParams{
			ID: 2,
		})

		if err != nil {
			return
		}

		rr2, _ := json.Marshal(resRoom2)
		var room2 core.RoomParams
		json.Unmarshal(rr2, &room2)

		resFromSCAM2, err := http.Get(urlSeyesCam + "/" + room2.UuidCam + "/channel" + "/0") //+ "/" + ps.Uuid + "/channel" + ps.Channel
		if err != nil {
			return
		}

		responseData2, err := ioutil.ReadAll(resFromSCAM2.Body)
		if err != nil {
			return
		}
		var dp2 core.DetectionParams
		json.Unmarshal(responseData2, &dp2)

		var jsonStr2 = []byte(`{"image":` + `"` + dp2.ImageData + `"` + `}`)
		req2, err := http.NewRequest("POST", "http://202.44.35.76:9094/detect", bytes.NewBuffer(jsonStr2))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client2 := &http.Client{}
		resp2, err := client2.Do(req2)
		if err != nil {
			return
		}
		defer resp2.Body.Close()

		body2, _ := ioutil.ReadAll(resp2.Body)
		json.Unmarshal(body2, &dp2)

		personCount2, _ := strconv.Atoi(dp2.PersonCount)
		comCount2, _ := strconv.Atoi(dp2.ConOnCount)

		resReports2, err := core.CreateReport(db, &core.ReportsParams{
			PersonCont: int64(personCount2),
			ComOnCount: int64(comCount2),
			Status:     "detected",
			Image:      dp.ImageData,
			Accurency:  dp.Accurency,
			RoomLabel:  room.Label,
			ReportTime: dp.Time,
			ReportDate: dp.Date,
		})

		if err != nil {
			return
		}

		rreport2, _ := json.Marshal(resReports2)
		var report2 core.ReportsParams
		json.Unmarshal(rreport2, &report2)

		notifyData2 := noti.NotifyParamV2{
			Image:     report2.Image,
			ID:        report2.ID,
			Person:    strconv.Itoa(int(report2.PersonCont)),
			ComOn:     strconv.Itoa(int(report2.ComOnCount)),
			UploadAt:  report2.ReportDate,
			Time:      report2.ReportTime,
			Accurency: report2.Accurency,
		}

		err = noti.SendToLineNotifyV2(&notifyData2)

		if err != nil {
			return
		}
	})

	logrus.Print(cornTime)
	// Start cron with one scheduled job
	logrus.Info("Start cron !")
	time.Sleep(5 * time.Second)

	c.Start()
}

func sanityChecks() error {
	_, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		return errors.New("Sanity check failure: " + err.Error())
	}

	return nil
}

func loadEnv() error {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		return errors.New("configuration_not_found")
	}

	logrus.Info("app env:", appEnv)

	return nil
}

func initSetting(db *gorm.DB) error {
	var setting model.Setting

	if err := db.First(&setting).Error; err != nil {
		if _, err := core.CreateSettings(db, &core.SettingsParams{
			AiModelData:           os.Getenv("DEFAULT_AI_MODEL"),
			CronjobTime:           os.Getenv("DEFAULT_CRONJOB_TIME"),
			LineNotifyAccessToken: os.Getenv("DEFAULT_LINE_NOTIFY_ACCESS_TOKEN"),
			MqttIp:                "",
			MqttUserName:          "",
			MqttPassword:          "",
			MqttPort:              "",
			MqttClientName:        "",
		}); err != nil {
			return err
		}
	}

	return nil
}
