package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
