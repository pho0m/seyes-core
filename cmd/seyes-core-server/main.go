package main

import (
	"errors"
	"fmt"
	"os"
	"seyes-core/internal/service"
	"seyes-core/internal/web"

	"time"
)

func main() {
	if err := loadEnv(); err != nil {
		panic(err)
	}

	sc, err := service.NewContainer()

	if err != nil {
		panic(err)
	}

	if err := sanityChecks(); err != nil {
		panic(err)
	}

	s := web.NewServer(sc, "3000")

	if err := service.DoMigration(sc.DB); err != nil {
		panic("cannot initialize Database: " + err.Error())
	}

	fmt.Println("Starting HTTP server...")
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
