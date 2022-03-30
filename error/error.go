package error

import (
	"net/http"
	"strings"
)

const InvalidCoordinates = "longitude and latitude should be in the right range " +
	"(-180<=longitude<=180 and -90<=latitude<=90)"

const UnprocessableCoordinates = "longitude and latitude should be number and not empty"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) AsResponse() *Error {
	return &Error{
		Message: e.Message,
	}
}

func (e Error) AsMessage() string {
	return e.Message
}

func NotFoundError(msg string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func ServerError(msg string) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func UnAuthorizedError(msg string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func ValidationError(msg string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func ParsingError(msg string) *Error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: strings.TrimSpace(msg),
	}
}
