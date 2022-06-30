package logger

type Fields map[string]interface{}

func Info(message string) {
	logger.Info(message)
}

func Debug(message string) {
	logger.Debug(message)
}

func Error(message string) {
	logger.Error(message)
}

func PrintAnything(msg interface{}) {
	logger.Error(msg)
}

func Panic(message string) {
	logger.Panic(message)
}
