package main

import (
	"errors"
	"fmt"
	"os"
	core "seyes-core/internal/core/dashboard"
	model "seyes-core/internal/model/room"
	"seyes-core/internal/service"
	"seyes-core/internal/web"

	"time"

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

	fmt.Println("Starting seyes http server...")
	fmt.Println("Listening in port:" + appPort)

	s.Start(sc)
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

	fmt.Println("app env:", appEnv)

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
		}); err != nil {
			return err
		}
	}

	return nil
}
