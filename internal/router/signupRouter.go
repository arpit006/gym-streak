package router

import (
	"github.com/arpit006/gym_streak/internal/handlers"
	"net/http"
)

type SignupRouter struct {
}

func (sr SignupRouter) InitializeRoutes() {
	router.HandleFunc("/signup", handlers.HandlePostSignupCall).Methods(http.MethodPost).Name("POST /signup")
}
