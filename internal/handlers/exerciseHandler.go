package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/arpit006/gym_streak/internal/exceptions"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/models"
	"github.com/arpit006/gym_streak/internal/response"
	"github.com/arpit006/gym_streak/internal/service"
	"net/http"
)

var exerciseService = service.GetExerciseServiceInstance()

func HandleAddExerciseCall(w http.ResponseWriter, r *http.Request) {
	logger.Info("Call to add exercise")

	var e models.Exercise
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		ex := exceptions.NewHttpException(http.StatusBadRequest, exceptions.BAD_REQUEST, fmt.Sprintf("Error in unmarshalling JSON. Error is %s", err), "bad request", map[string]string{
			"api": "POST /exercise",
		})
		response.HandleErrorResponse(w, ex)
		return
	}
	logger.Infof("Exercise Object is %+v", e)

	ex := exerciseService.AddAnExercise(&e)

	if ex != nil {
		response.HandleErrorResponse(w, ex)
		return
	}
	bodyJson, jsonErr := json.Marshal(e)
	if jsonErr != nil {
		ex := exceptions.NewHttpException(http.StatusInternalServerError, exceptions.INTERNAL_SERVER_ERROR, fmt.Sprintf("Error in marshalling Exercise body to JSON. Error is %s", err), "internal server error", map[string]string{
			"api": "POST /exercise",
		})
		response.HandleErrorResponse(w, ex)
		return
	}
	resp := response.GymStreakResp{
		HttpStatusCode: http.StatusCreated,
		Msg:            string(bodyJson),
		Headers: map[string]string{
			"api": "POST /exercise",
		},
	}
	response.GenerateResponse(w, resp)
}
