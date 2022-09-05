package router

import (
	"github.com/arpit006/gym_streak/internal/handlers"
	"net/http"
)

type ExerciseRouter struct{}

func (e ExerciseRouter) InitializeRoutes() {
	router.HandleFunc("/exercise", handlers.HandleAddExerciseCall).Methods(http.MethodPost).Name("POST /exercise")
}
