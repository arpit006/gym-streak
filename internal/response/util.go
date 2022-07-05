package response

import (
	"github.com/arpit006/gym_streak/internal/exceptions"
	"github.com/arpit006/gym_streak/internal/logger"
	"net/http"
)

func GenerateResponse(w http.ResponseWriter, resp GymStreakResp) error {
	logger.PrintAnything("Printing the response to be created ")
	logger.PrintAnything(resp)

	// writing headers to the response
	for k, v := range resp.Headers {
		w.Header().Set(k, v)
	}

	// TODO: add middleware for these tasks
	// adding custom Header info
	w.Header().Set("developer-email", "iamarpitsrivastava06@gmail.com")
	w.Header().Set("company", "Gym-Streak-App")
	w.Header().Set("version", "0.0.1")

	//	writing response
	_, err := w.Write([]byte(resp.Msg))
	if err != nil {
		logger.Error("Error occurred while preparing response")
		return exceptions.ThrowHttpResponseException()
	}
	logger.Info("Message written successfully to the response")

	//	 set response code
	w.WriteHeader(resp.HttpStatusCode)

	return nil
}

func GenerateErrorResponse(w http.ResponseWriter, err error) {

}
