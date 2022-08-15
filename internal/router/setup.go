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
	var testRouter GymStreakRouters = PingRouter{}
	testRouter.InitializeRoutes()

	var signupRouter GymStreakRouters = SignupRouter{}
	signupRouter.InitializeRoutes()

	return router
}
