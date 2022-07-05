package logger

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/config"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger *logrus.Logger = nil

// Init Singleton instance of logger
func Init() {
	if logger != nil {
		return
	}
	loggerConfig := config.GetLoggingConfig()
	logLevel, err := logrus.ParseLevel(loggerConfig.Level)
	if err != nil {
		panic(fmt.Sprintf("Not supported logging level: [%s]", loggerConfig.Level))
	}
	logFormatter := loggingFormatterFactory(loggerConfig.Formatter)
	logger = logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(logLevel)
	logger.SetFormatter(logFormatter)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/gym-streak-app.log",
		MaxSize:    5, // MB
		MaxBackups: 10,
		MaxAge:     30, // days
		Compress:   true,
	}
	logger.SetOutput(io.MultiWriter(lumberjackLogger, os.Stdout))
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
