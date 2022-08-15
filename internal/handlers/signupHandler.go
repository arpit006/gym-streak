package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/arpit006/gym_streak/internal/constants"
	"github.com/arpit006/gym_streak/internal/db"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/models"
	"github.com/arpit006/gym_streak/internal/response"
	"github.com/arpit006/gym_streak/internal/stores"
	"net/http"
)

func HandleSignupPostCalls(w http.ResponseWriter, r *http.Request) {
	logger.Info("Signup call request received..")

	var u models.User

	postBody := r.Body
	err := json.NewDecoder(postBody).Decode(&u)
	appVersion := r.Header.Get(constants.APP_VERSION)

	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signupStore := stores.NewSignupStore(db.GetDB())
	signupStore.Save(u)

	respBody := response.GymStreakResp{
		HttpStatusCode: http.StatusOK,
		Msg:            "POST /signup method is working absolutely fine!",
		Headers: map[string]string{
			"api":        "POST /signup",
			"appVersion": appVersion,
		},
	}
	err = response.GenerateResponse(w, respBody)
	logger.Error(fmt.Sprintf("%s", err))
}
