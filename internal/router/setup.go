package router

import (
	"github.com/gorilla/mux"
)

var router *mux.Router = nil

// initRouterInstance singleton implementation of router
func initRouterInstance() *mux.Router {
	if router != nil {
		return router
	}
	router = mux.NewRouter()
	return router
}

func InitRoutes() *mux.Router {
	initRouterInstance()
	var pingRouter GymStreakRouters = PingRouter{}
	pingRouter.InitializeRoutes()

	var exerciseRouter GymStreakRouters = ExerciseRouter{}
	exerciseRouter.InitializeRoutes()

	return router
}
