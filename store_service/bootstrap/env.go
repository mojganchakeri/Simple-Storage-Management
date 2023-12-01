package bootstrap

import (
	"store_service/internal/models"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupEnv() (env models.Env) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		logrus.Fatal("Environment can't be loaded: ", err)
	}

	return env
}
