package bootstrap

import (
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func SetupLogger(serviceName string, logLevel string) error {
	file, err := os.OpenFile(fmt.Sprintf("%s.log", serviceName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)

	switch strings.ToLower(logLevel) {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "debug":
		log.SetLevel(log.FatalLevel)
	default:
		return errors.New("log level is not valid")
	}

	return nil

}
