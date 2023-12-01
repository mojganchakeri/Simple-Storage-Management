package repository

import (
	"store_service/configs"
	"store_service/internal/models"
)

func (dbClient *Repository) GetTagId(value string) (string, error) {
	var tags []models.TagGorm

	res := dbClient.DB.Table(configs.TagTable).Where("value=?", value).Scan(&tags)
	if res.Error != nil {
		return "", res.Error
	} else if len(tags) != 0 {
		return tags[0].ID, nil
	} else {
		return "", nil
	}
}
