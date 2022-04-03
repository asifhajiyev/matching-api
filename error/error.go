package error

import (
	"net/http"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type FieldValidationError struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
}

func ServerError(details interface{}) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Details: details,
	}
}

func UnAuthorizedError(m string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: m,
	}
}

func ParsingError(details interface{}) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Details: details,
	}
}

func ValidationError(details interface{}) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Details: details,
	}
}
