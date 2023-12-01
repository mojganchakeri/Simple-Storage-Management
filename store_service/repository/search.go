package repository

import (
	"fmt"
	"store_service/configs"
	"store_service/internal/models"

	"github.com/sirupsen/logrus"
)

func (dbClient *Repository) SearchFileInDB(name string, tags []string) ([]string, error) {
	searchFiles := []models.FileGorm{}
	searchFileNames := []string{}

	searchTx := dbClient.DB.Table(configs.StoreTable)

	if name != "" {
		searchTx.Where("name=?", name)
	}

	if len(tags) != 0 {
		joinQuery := fmt.Sprintf("left join %s on %s.id = %s.file_id", configs.StoreTagTable, configs.StoreTable, configs.StoreTagTable)
		searchTx = searchTx.Joins(joinQuery)
		joinQuery = fmt.Sprintf("left join %s on %s.tag_id = %s.id", configs.TagTable, configs.StoreTagTable, configs.TagTable)
		searchTx = searchTx.Joins(joinQuery)

		searchTx.Where("tag.value in ?", tags)
	}

	res := searchTx.Scan(&searchFiles)

	if res.Error != nil {
		logrus.Error("search query fail")
		return searchFileNames, res.Error
	}

	if len(searchFiles) == 0 {
		dbClient.DB.Table(configs.StoreTable).Order("created_at asc").Limit(1).Scan(&searchFiles)
	}

	for _, res := range searchFiles {
		searchFileNames = append(searchFileNames, res.FilePath)
	}
	return searchFileNames, nil

}
