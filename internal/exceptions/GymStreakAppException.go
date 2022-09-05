package exceptions

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	DEFAULT_ERROR_CODE          = 500
	DEFAULT_ERROR_MESSAGE       = "Internal Server Error"
	DEFAULT_SHORT_ERROR_MESSAGE = "internal server error"
	DEFAULT_ERROR_TYPE          = ErrorType("internal-server-error")
)

// ErrorType string
type ErrorType string

var (
	BAD_REQUEST           = mapToErrorType("bad-request")
	INTERNAL_SERVER_ERROR = mapToErrorType("internal-server-error")
	DATABASE_SAVE_ERROR   = mapToErrorType("database-save-error")
)

type GymStreakAppException interface {
	GetErrorCode() int              //400
	GetErrorMessage() string        // invalid request body
	GetShortErrorMessage() string   // request body in incorrect
	GetErrorDisplayMessage() string // [code, type, shortMsg]
	GetErrorType() ErrorType        // internal-server-error
	GetErrorHeaders() map[string]string
	Error() string
}

func getDefaultError() error {
	//TODO: throw a default error here
	return errors.New(fmt.Sprintf(""))
}

func mapToErrorType(s string) ErrorType {
	return ErrorType(s)
}

func (et ErrorType) GetErrorTypeValue() string {
	return string(et)
}

func getDefaultErrorDisplayMessage() string {
	body := map[string]string{
		"code": string(rune(DEFAULT_ERROR_CODE)),
		"type": string(DEFAULT_ERROR_TYPE),
		"msg":  DEFAULT_SHORT_ERROR_MESSAGE,
	}
	jsonString, _ := json.Marshal(body)
	return string(jsonString)
}
