package errors

import "net/http"

type Http struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHttp(code int, message string) error {
	return &Http{
		Code:    code,
		Message: message,
	}
}

func (e Http) Error() string {
	return e.Message
}

func HttpBadRequest(message string) error {
	return NewHttp(http.StatusBadRequest, message)
}
