package handlers

import (
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/response"
	"net/http"
)

func HandleTestGetMethod(w http.ResponseWriter, r *http.Request) {
	logger.Info("Inside Test Handler Get method")

	respBody := response.GymStreakResp{
		HttpStatusCode: http.StatusOK,
		Msg:            "/test method is working absolutely fine!",
		Headers: map[string]string{
			"api": "/test",
		},
	}
	err := response.GenerateResponse(w, respBody)
	if err != nil {
		logger.PrintAnything(err)
		return
	}
}
