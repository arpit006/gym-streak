package main

import (
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/setup"
)

func main() {
	// base setup
	setup.Setup()

	logger.Info("Gym-Streak application started..........")
	logger.Debug("Gym-Streak application started..........")
	logger.Error("Gym-Streak application started..........")
}
