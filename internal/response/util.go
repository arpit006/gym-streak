package response

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/exceptions"
	"github.com/arpit006/gym_streak/internal/logger"
	"net/http"
)

func GenerateResponse(w http.ResponseWriter, resp GymStreakResp) {
	logger.Info("Printing the response to be created ")
	logger.Info(fmt.Sprintf("%+v", resp))

	// writing headers to the response
	if resp.Headers != nil {
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
	}

	// TODO: add middleware for these tasks
	// adding custom Header info
	w.Header().Set("developer-email", "arpittech06@gmail.com")
	w.Header().Set("company", "Gym-Streak-App")
	w.Header().Set("version", "0.0.1")

	//	 set http-response code
	w.WriteHeader(resp.HttpStatusCode)

	//	writing response
	_, err := w.Write([]byte(resp.Msg))
	if err != nil {
		logger.Error("Error occurred while preparing response")
		// TODO: here
		return
	}
	logger.Info("Message written successfully to the response")
}

func HandleErrorResponse(w http.ResponseWriter, err exceptions.GymStreakAppException) {
	if err == nil {
		//	Handle internal server error here
	}
	logger.Error(err.Error())
	headers := err.GetErrorHeaders()

	if headers != nil {
		for k, v := range headers {
			w.Header().Set(k, v)
		}
	}

	// TODO: add middleware for these tasks
	// adding custom Header info
	w.Header().Set("developer-email", "arpittech06@gmail.com")
	w.Header().Set("company", "Gym-Streak-App")
	w.Header().Set("version", "0.0.1")

	//	 set http-response code
	w.WriteHeader(err.GetErrorCode())

	//	writing response
	_, e := w.Write([]byte(err.GetErrorDisplayMessage()))
	if e != nil {
		logger.Errorf("Error occurred while writing error response.. error is %s", e.Error())
		return
	}
	logger.Info("Error response written successfully to the response........")
}
