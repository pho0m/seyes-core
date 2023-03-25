package main

import (
	"encoding/json"
	"errors"
	"os"
	core "seyes-core/internal/core/dashboard"
	"seyes-core/internal/helper"
	model "seyes-core/internal/model/room"
	"seyes-core/internal/service"
	"seyes-core/internal/web"

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

	var cornTime = ""

	cornTime = os.Getenv("DEFAULT_CRONJOB_TIME")
	c := cron.New()

	sdata, err := core.GetSetting(db, &helper.UrlParams{ID: 1})

	if err != nil {
		panic(err)
	}
	sm, _ := json.Marshal(sdata)
	var setting core.SettingsParams
	json.Unmarshal(sm, &setting)

	cornTime = setting.CronjobTime

	c.AddFunc("@every "+cornTime+"s", func() {

		logrus.Info("[Job] Now Every " + cornTime + " job")

	})
	logrus.Print(cornTime)
	// Start cron with one scheduled job
	logrus.Info("Start cron !")
	time.Sleep(2 * time.Second)

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
