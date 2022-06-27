package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func Init() {
	// TODO: get this config from viper/vault
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
}
