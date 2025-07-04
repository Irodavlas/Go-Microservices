package model

import (
	"log"

	"github.com/labstack/echo/v4"
)

type SuccessGetResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    User   `json:"data,omitempty"`
}
type FailedGetResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// define other method for succesfull responses and not
func Success(c echo.Context, statusCode int, message string, user User) error {
	log.Printf("[INFO] CODE:%d, MESSAGE:%s", statusCode, message)
	return c.JSON(statusCode, SuccessGetResponse{
		Code:    statusCode,
		Success: true,
		Message: message,
		Data:    user,
	})
}
func Failed(c echo.Context, statusCode int, message string, err string) error {
	log.Printf("[INFO] CODE:%d, MESSAGE:%s", statusCode, message)
	return c.JSON(statusCode, FailedGetResponse{
		Code:    statusCode,
		Success: false,
		Message: message,
		Error:   err,
	})
}
