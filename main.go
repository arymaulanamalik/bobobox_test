package main

import (
	"log"

	"github.com/arymaulanamalik/bobobox_test/config"
	_ "github.com/arymaulanamalik/bobobox_test/pkg/datadog"
	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
	"github.com/arymaulanamalik/bobobox_test/server"
)

func main() {
	var (
		port = config.API_PORT
	)

	logger.Info("Service Port : ", port)

	runner, err := server.NewRun()
	if err != nil {
		log.Println(err.Error())
	}

	if err := runner.Run(port); err != nil {
		log.Println(err.Error())
	}
}
