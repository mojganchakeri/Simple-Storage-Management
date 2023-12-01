package repository

import (
	"retreival_service/internal/models"

	"github.com/sirupsen/logrus"
)

func (dbClient *Repository) Migrate() {
	err := dbClient.DB.AutoMigrate(
		&models.SchemaUser{},
	)
	if err != nil {
		logrus.Error(err)
	}
}
