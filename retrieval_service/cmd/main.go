package main

import (
	"retreival_service/api"
	"retreival_service/bootstrap"
	"retreival_service/configs"
	"retreival_service/repository"
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
