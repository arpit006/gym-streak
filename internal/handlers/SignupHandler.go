package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/arpit006/gym_streak/internal/constants"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/models"
	"github.com/arpit006/gym_streak/internal/response"
	"github.com/arpit006/gym_streak/internal/stores"
	"net/http"
)

// SignupStore dependency
var signupStore *stores.SignupStore = stores.GetSignupStoreInstance()

func HandlePostSignupCall(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("inside Handle Post Signup call."))

	postBody := r.Body
	appVersionHeader := r.Header.Get(constants.APP_VERSION_HEADER)

	var u models.User

	err := json.NewDecoder(postBody).Decode(&u)
	if err != nil {
		logger.Error(fmt.Sprintf("Error in unmarshalling JSON. Error is %s", err))
	}

	// save signup info to DB
	err = signupStore.Signup(u)
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
		errBody := response.GymStreakResp{
			HttpStatusCode: http.StatusInternalServerError,
			Msg:            fmt.Sprintf("error occurred in POST /signup call. Error is %s", err),
			Headers: map[string]string{
				"api":        "POST /signup",
				"appVersion": appVersionHeader,
			},
		}
		err = response.GenerateResponse(w, errBody)
		return
	}

	respBody := response.GymStreakResp{
		HttpStatusCode: http.StatusOK,
		Msg:            "/signup method is working absolutely fine!",
		Headers: map[string]string{
			"api":        "/test",
			"appVersion": appVersionHeader,
		},
	}
	err = response.GenerateResponse(w, respBody)
	if err != nil {
		logger.PrintAnything(err)
		return
	}

}
