package router

import (
	"github.com/arpit006/gym_streak/internal/handlers"
	"net/http"
)

type TestRouter struct{}

func (t TestRouter) InitializeRoutes() {
	router.HandleFunc("/test", handlers.HandleTestGetMethod).Methods(http.MethodGet)
}
