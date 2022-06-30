package setup

import (
	"github.com/arpit006/gym_streak/internal/config"
	"github.com/arpit006/gym_streak/internal/logger"
)

func Setup() {
	//	initialize config
	config.Load()
	//	setup logging
	setupLogging()
	//config.Load()
}

func setupLogging() {
	logger.Init()
}
