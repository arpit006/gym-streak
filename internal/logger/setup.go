package logger

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/config"
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func Init() {
	// TODO: get this config from viper/vault
	loggerConfig := config.GetLoggingConfig()
	logLevel, err := logrus.ParseLevel(loggerConfig.Level)
	if err != nil {
		panic(fmt.Sprintf("Not supported logging level: [%s]", loggerConfig.Level))
	}
	logFormatter := loggingFormatterFactory(loggerConfig.Formatter)
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetLevel(logLevel)
	logger.SetFormatter(logFormatter)
}

func loggingFormatterFactory(formatter string) logrus.Formatter {
	switch formatter {
	case "json":
		return &logrus.JSONFormatter{}
	case "text":
		return &logrus.TextFormatter{}
	default:
		panic(fmt.Sprintf("Not supported logging Formatting: [%s]", formatter))

	}
}
