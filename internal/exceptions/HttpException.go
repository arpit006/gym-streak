package exceptions

import (
	"encoding/json"
	"fmt"
)

type HttpException struct {
	Code     int
	Msg      string
	ShortMsg string
	Type     ErrorType
	Headers  map[string]string
}

func NewHttpException(code int, errorType ErrorType, msg string, shortMsg string, headers map[string]string) HttpException {
	return HttpException{
		Code:     code,
		Msg:      msg,
		ShortMsg: shortMsg,
		Type:     errorType,
		Headers:  headers,
	}
}

func (ex HttpException) GetErrorCode() int {
	return ex.Code
}

func (ex HttpException) GetErrorMessage() string {
	return ex.Msg
}
func (ex HttpException) GetErrorType() ErrorType {
	return ex.Type
}

func (ex HttpException) GetErrorHeaders() map[string]string {
	return ex.Headers
}

func (ex HttpException) GetErrorDisplayMessage() string {
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

func (ex HttpException) GetShortErrorMessage() string {
	return ex.ShortMsg
}

func (ex HttpException) Error() string {
	return fmt.Sprintf("ERROR : [%s] [%d] [%s] [%s]", ex.GetErrorType().GetErrorTypeValue(), ex.GetErrorCode(), ex.GetShortErrorMessage(), ex.GetErrorMessage())
}
