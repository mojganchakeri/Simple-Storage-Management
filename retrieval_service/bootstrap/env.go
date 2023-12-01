package bootstrap

import (
	"retreival_service/internal/models"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupEnv() (env models.Env) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return env
}
