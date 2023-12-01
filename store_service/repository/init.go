package repository

import (
	"fmt"
	"store_service/bootstrap"
	"store_service/configs"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

var DBClient = &Repository{}

func SetClient() {
	log.Info(fmt.Sprintf("connecting to database - host: %s, port: %s, dbname: %s", configs.Env.DBHost, configs.Env.DBPort, configs.Env.DBName))
	dbClient := bootstrap.ConnectDatabaseMariadb(
		configs.Env.DBUser,
		configs.Env.DBPass,
		configs.Env.DBHost,
		configs.Env.DBPort,
		configs.Env.DBName,
	)
	DBClient.DB = dbClient.DB
	DBClient.Migrate()
}
