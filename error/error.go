package error

import (
	"net/http"
	"strings"
)

const InvalidCoordinates = "longitude and latitude should be in the right range " +
	"(-180<=longitude<=180 and -90<=latitude<=90)"

const UnprocessableCoordinates = "longitude and latitude should be number and not empty"

const URLNotFound = "requested url does not exist"

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

func ServerError(m string) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: m,
	}
}

func UnAuthorizedError(m string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: m,
	}
}

func ValidationError(m string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: m,
	}
}

func ParsingError(m string) *Error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: strings.TrimSpace(m),
	}
}
