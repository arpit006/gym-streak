package config

type LoggerConfig struct {
	Level     string
	Formatter string
}

func readLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Level:     readViperString("logging_level"),
		Formatter: readViperString("logging_formatter"),
	}
}
