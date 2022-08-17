package handlers

import (
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/response"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	logger.InfoF("Inside Test Handler Get method")

	respBody := response.GymStreakResp{
		HttpStatusCode: http.StatusOK,
		Msg: "{\n	\"pong\"\n}",
		Headers: map[string]string{
			"api": "GET /ping",
		},
	}
	err := response.GenerateResponse(w, respBody)
	if err != nil {
		logger.ErrorF(err.Error())
		return
	}
}
