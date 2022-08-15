package router

import (
	"github.com/arpit006/gym_streak/internal/handlers"
	"net/http"
)

type PingRouter struct{}

func (p PingRouter) InitializeRoutes() {
	router.HandleFunc("/ping", handlers.HandlePing).Methods(http.MethodGet)
}
