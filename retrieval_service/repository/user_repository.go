package repository

import (
	"retreival_service/configs"
	"retreival_service/internal/models"

	"github.com/sirupsen/logrus"
)

func (dbClient *Repository) GetUserByUsername(username string) (models.User, error) {
	var users []models.User

	res := dbClient.DB.Table(configs.UserTable).Where("user.username=?", username).Scan(&users)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	if len(users) == 0 {
		return models.User{}, nil
	}

	return users[0], nil

}

func (dbClient *Repository) GetUserByID(id string) (models.User, error) {
	var users []models.User

	res := dbClient.DB.Table(configs.UserTable).Where("user.id=?", id).Scan(&users)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	if len(users) == 0 {
		return models.User{}, nil
	}

	return users[0], nil
}

func (dbClient *Repository) AddNewUser(user models.User) error {
	if err := dbClient.DB.Table(configs.UserTable).Create(&user).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
