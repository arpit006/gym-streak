package setup

import (
	"github.com/arpit006/gym_streak/internal/config"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/router"
	"github.com/arpit006/gym_streak/internal/stores"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Setup() {
	//	initialize config
	config.Load()
	//	setup logging
	setupLogging()
	// setup database
	setupDatabase()
	// setup tables
	setupTables()
	// setup router
	r := setupRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//logger.InfoF("Gym-Streak application started..........")
	//logger.DebugF("Gym-Streak application started..........")
	//logger.ErrorF("Gym-Streak application started..........")

	logger.InfoF("arpit", "srivastava")

	log.Fatal(srv.ListenAndServe())
}

func setupLogging() {
	logger.Init()
}

func setupRouter() *mux.Router {
	return router.InitRoutes()
}

func setupTables() {
	err := stores.RunDatabaseMigrations()
	if err != nil {
		logger.PanicF(err.Error())
	}
}

func setupDatabase() {
	stores.Init()
}
