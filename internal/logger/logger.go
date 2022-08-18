package logger

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/constants"
	"runtime"
	"strings"
)

const (
	skipCallers = 3
)

type Fields map[string]interface{}

func getSourceInfo() map[string]interface{} {
	file, line := getFileInfo()
	return map[string]interface{}{
		"file-source": fmt.Sprintf("%s:%d", file, line),
	}
}

func getFileInfo() (string, int) {
	_, file, line, _ := runtime.Caller(skipCallers)
	return chopPath(file), line
}

func chopPath(filePath string) string {
	i := strings.Index(filePath, constants.APP_NAME)
	if i != -1 {
		return filePath[i:]
	}
	return filePath
}

func Infof(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Infof(msg, args)
}

func Errorf(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Errorf(msg, args)
}

func Debugf(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Debugf(msg, args)
}

func Warnf(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Warnf(msg, args)
}

func Panicf(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Panicf(msg, args)
}

func Fatalf(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Fatalf(msg, args)
}

func Info(msg string) {
	logger.WithFields(getSourceInfo()).Info(msg)
}

func Warn(msg string) {
	logger.WithFields(getSourceInfo()).Warn(msg)
}

func Error(msg string) {
	logger.WithFields(getSourceInfo()).Error(msg)
}

func Debug(msg string) {
	logger.WithFields(getSourceInfo()).Debug(msg)
}

func Panic(msg string) {
	logger.WithFields(getSourceInfo()).Panic(msg)
}

func Fatal(msg string) {
	logger.WithFields(getSourceInfo()).Fatal(msg)
}
