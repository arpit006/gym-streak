package router

import (
	"github.com/arpit006/gym_streak/internal/handlers"
	"net/http"
)

type Test struct{}

func (t Test) InitializeRoutes() {
	router.HandleFunc("/test", handlers.HandleTestGetMethod).Methods(http.MethodGet)
}
