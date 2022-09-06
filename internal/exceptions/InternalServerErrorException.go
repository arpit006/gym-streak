package exceptions

import (
	"encoding/json"
	"fmt"
)

type InternalServerError struct {
	Code     int
	Msg      string
	ShortMsg string
	Type     ErrorType
	Headers  map[string]string
}

func NewInternalServerError() InternalServerError {
	return InternalServerError{
		Code:     500,
		Msg:      "Internal Server Error",
		ShortMsg: "internal server error",
		Type:     INTERNAL_SERVER_ERROR,
		Headers: map[string]string{
			"error": "internal-server-error",
			"type":  "unhandled exception",
		},
	}
}

func (ex InternalServerError) GetErrorCode() int {
	return ex.Code
}

func (ex InternalServerError) GetErrorMessage() string {
	return ex.Msg
}
func (ex InternalServerError) GetErrorType() ErrorType {
	return ex.Type
}

func (ex InternalServerError) GetErrorHeaders() map[string]string {
	return ex.Headers
}

func (ex InternalServerError) GetErrorDisplayMessage() string {
	body := map[string]interface{}{
		"code": ex.GetErrorCode(),
		"type": ex.GetErrorType(),
		"msg":  ex.GetShortErrorMessage(),
	}
	jsonString, err := json.Marshal(body)
	if err != nil {
		getDefaultErrorDisplayMessage()
	}
	return string(jsonString)
}

func (ex InternalServerError) GetShortErrorMessage() string {
	return ex.ShortMsg
}

func (ex InternalServerError) Error() string {
	return fmt.Sprintf("ERROR : [%s] [%d] [%s] [%s]", ex.GetErrorType().GetErrorTypeValue(), ex.GetErrorCode(), ex.GetShortErrorMessage(), ex.GetErrorMessage())
}
