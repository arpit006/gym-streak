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

func InfoF(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Infof(msg, args)
}

func ErrorF(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Errorf(msg, args)
}

func DebugF(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Debugf(msg, args)
}

func WarnF(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Warnf(msg, args)
}

func PanicF(msg string, args ...interface{}) {
	logger.WithFields(getSourceInfo()).Panicf(msg, args)
}
