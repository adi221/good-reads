package helper

import "net/http"

type HttpError struct {
	Code       string `json:"code"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Error is implementing Error interface
func (httpErr HttpError) Error() string {
	return httpErr.Message
}

// Extensions allow to add extra info to graphqlerr
func (httpErr HttpError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"message":    httpErr.Message,
		"code":       httpErr.Code,
		"statusCode": httpErr.StatusCode,
	}
}

func NewHttpError(code string, statusCode int, message string) HttpError {
	return HttpError{Code: code, StatusCode: statusCode, Message: message}
}

// Common http error codes
const (
	InternalServerError = "ERR_INTERNAL_SERVER_ERROR"
)

// ServerError means that a certain server error occurred.
func ServerError() HttpError {
	return NewHttpError(InternalServerError, http.StatusInternalServerError, "Internal Server Error")
}
