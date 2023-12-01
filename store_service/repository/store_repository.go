package repository

import (
	"store_service/configs"
	"store_service/internal/models"

	"gorm.io/gorm"
)

func (dbClient *Repository) StoreFile(fileObj models.FileGorm, tagsObj []models.TagGorm, fileTagObj []models.FileTagGorm) error {
	return dbClient.DB.Transaction(func(tx *gorm.DB) error {

		if fileObj.Name != "" {
			if err := tx.Table(configs.StoreTable).Create(&fileObj).Error; err != nil {
				return err
			}
		}

		if len(tagsObj) != 0 {
			if err := tx.Table(configs.TagTable).Create(&tagsObj).Error; err != nil {
				return err
			}
		}

		if len(fileTagObj) != 0 {
			if err := tx.Table(configs.StoreTagTable).Create(&fileTagObj).Error; err != nil {
				return err
			}
		}

		return nil
	})

}

func (dbClient *Repository) DeleteFile(filesPath []string) error {

	res := dbClient.DB.Table(configs.StoreTable).Where("file_path in ?", filesPath).Delete(&models.FileGorm{})
	return res.Error

}
