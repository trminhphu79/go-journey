package network

import (
	"errors"
	"fmt"
	"net/http"
)

type apiError struct {
	Code    int
	Message string
	Err     error
}

func (e *apiError) GetCode() int {
	return e.Code
}

func (e *apiError) GetMessage() string {
	return e.Message
}

func (e *apiError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%d - %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func (e *apiError) Unwrap() error {
	return e.Err
}

func newApiError(code int, message string, err error) ApiError {
	apiError := apiError{
		Code:    code,
		Message: message,
		Err:     err,
	}
	if err == nil {
		apiError.Err = errors.New(message)
	}
	return &apiError
}
func NewBadRequestErr(message string, err error) ApiError {
	return newApiError(http.StatusBadRequest, message, err)
}

func NewForbiddenErr(message string, err error) ApiError {
	return newApiError(http.StatusForbidden, message, err)
}

func NewUnauthorizedErr(message string, err error) ApiError {
	return newApiError(http.StatusUnauthorized, message, err)
}

func NewNotFoundErr(message string, err error) ApiError {
	return newApiError(http.StatusNotFound, message, err)
}

func NewInternalServerErr(message string, err error) ApiError {
	return newApiError(http.StatusInternalServerError, message, err)
}

func ErrorResponse(message string, err error) ApiError {
	return newApiError(http.StatusNotFound, message, err)
}
