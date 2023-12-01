package main

import (
	"fmt"
	"store_service/api"
	"store_service/bootstrap"
	"store_service/configs"
	repository "store_service/repository"

	"github.com/sirupsen/logrus"
)

func main() {

	// 1. Read configs from env
	configs.Env = bootstrap.SetupEnv()

	// 2. Setup logger
	bootstrap.SetupLogger(configs.ServiceName, configs.Env.LogLevel)

	// 3. Connection to DB
	repository.SetClient()

	// 4. Setup server
	server := api.SetupServer()

	logrus.Info(fmt.Sprintf("service starts at %s", configs.Env.ServerAddress))
	server.Run(configs.Env.ServerAddress)
}
