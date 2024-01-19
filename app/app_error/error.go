package app_error

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

const genericDetailErrorTemplate = "%s: %s"

type BaseError struct {
	UserMessage string
	Err         error
	HttpCode    int
}

type Error struct {
	BaseError
}

func NewError(err error, httpCode int, userMessage string) Error {
	return Error{
		BaseError{
			UserMessage: userMessage,
			Err:         err,
			HttpCode:    httpCode,
		},
	}
}

func GetGenericErrorMessageByHttpCode(httpStatusCode int) string {
	message := ""

	if httpStatusCode == 401 {
		return "unauthorized"
	} else if httpStatusCode == 400 {
		return "bad request"
	} else if httpStatusCode >= 500 && httpStatusCode <= 599 {
		return "internal server error"
	}

	return message
}

func (e Error) Error() string {
	return formatUserErrorMessage(GetGenericErrorMessageByHttpCode(e.HttpCode), e.UserMessage)
}

func (e Error) Log() Error {
	log.Printf(e.Error())
	log.Println(e.Err)
	log.Printf(string(debug.Stack()))

	return e
}

func (e Error) HttpError(w http.ResponseWriter) {
	http.Error(w, e.Error(), e.HttpCode)
}

func formatUserErrorMessage(genericMessage, userMessage string) string {
	if userMessage == "" {
		return genericMessage
	}

	return fmt.Sprintf(genericDetailErrorTemplate, genericMessage, userMessage)
}
