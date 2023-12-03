package bootstrap

import (
	"store_service/internal/models"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupEnv() (env models.Env) {
	v := viper.New()
	v.SetConfigFile("/store-service.yaml")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatal("Can't find the file yaml : ", err)
	}
	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	if err = v.MergeInConfig(); err != nil {
		logrus.Fatal("Environment can't be loaded: ", err)
	}
	err = v.Unmarshal(&env)
	if err != nil {
		logrus.Fatal("Environment can't be loaded: ", err)
	}

	logrus.Infof("config loaded: %+v", env)

	return env
}
