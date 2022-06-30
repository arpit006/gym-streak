package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type appConfig struct {
	logConfig LoggerConfig
	Rand      string
}

var conf appConfig

// Load the configuration from the Viper
func Load() {
	setupViperConfig()

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	conf = appConfig{
		logConfig: readLoggerConfig(),
	}
	//logger.PrintAnything(conf)
}

func setupViperConfig() {
	viper.SetConfigName("application-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	viper.AutomaticEnv()
}

func GetLoggingConfig() LoggerConfig {
	return conf.logConfig
}
