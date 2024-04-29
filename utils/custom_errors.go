package utils

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code:%d error:%s", c.StatusCode, c.Message)
}

func NewCustomError(code int, message string) error {
	return &CustomError{
		StatusCode: code,
		Message:    message,
	}
}

func NewBadRequestError(message string) error {
	return NewCustomError(http.StatusBadRequest, message)
}

func NewInternalServerError(message string) error {
	return &CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
