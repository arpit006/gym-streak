package exceptions

import (
	"encoding/json"
	"fmt"
)

type DatabaseException struct {
	Code     int
	Msg      string
	ShortMsg string
	Type     ErrorType
	Headers  map[string]string
}

func NewDatabaseException(code int, errorType ErrorType, msg string, shortMsg string, headers map[string]string) DatabaseException {
	return DatabaseException{
		Code:     code,
		Msg:      msg,
		ShortMsg: shortMsg,
		Type:     errorType,
		Headers:  headers,
	}
}

func (ex DatabaseException) GetErrorCode() int {
	return ex.Code
}

func (ex DatabaseException) GetErrorMessage() string {
	return ex.Msg
}
func (ex DatabaseException) GetErrorType() ErrorType {
	return ex.Type
}

func (ex DatabaseException) GetErrorHeaders() map[string]string {
	return ex.Headers
}

func (ex DatabaseException) GetErrorDisplayMessage() string {
	body := map[string]string{
		"code": string(rune(ex.GetErrorCode())),
		"type": string(ex.GetErrorType()),
		"msg":  ex.GetShortErrorMessage(),
	}
	jsonString, err := json.Marshal(body)
	if err != nil {
		getDefaultErrorDisplayMessage()
	}
	return string(jsonString)
}

func (ex DatabaseException) GetShortErrorMessage() string {
	return ex.ShortMsg
}

func (ex DatabaseException) Error() string {
	return fmt.Sprintf("ERROR : [%s] [%d] [%s] [%s]", ex.GetErrorType().GetErrorTypeValue(), ex.GetErrorCode(), ex.GetShortErrorMessage(), ex.GetErrorMessage())
}
