package main

import (
	"store_service/api"
	"store_service/bootstrap"
	"store_service/configs"
	repository "store_service/repository"
)

func main() {

	// 1. Read configs from env
	configs.Env = bootstrap.SetupEnv()

	// 2. Setup logger
	bootstrap.SetupLogger(configs.ServiceName, configs.Env.LogLevel)

	// 3. Connection to DB
	repository.SetClient()

	// 4. Setup server
	api.SetupServer()

}
