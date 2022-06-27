package setup

import "github.com/arpit006/gym_streak/internal/logger"

func Setup() {
	//	initialize config
	//	setup logging
	setupLogging()
}

func setupLogging() {
	logger.Init()
}
