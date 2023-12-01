package repository

import "store_service/internal/models"

func (dbClient *Repository) Migrate() {
	dbClient.DB.AutoMigrate(
		&models.SchemaStorage{},
		&models.SchemaTag{},
		&models.SchemaStoreTag{},
	)
}
