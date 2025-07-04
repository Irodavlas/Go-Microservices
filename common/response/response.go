package response

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int          `json:"code"`
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Error   string       `json:"error,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}
type Sender struct{}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Success(c echo.Context, statusCode int, message string, data interface{}) error {
	log.Printf("[INFO] CODE:%d, MESSAGE:%s", statusCode, message)
	return c.JSON(statusCode, Response{
		Code:    statusCode,
		Success: true,
		Message: message,
		Data:    &data,
	})
}
func (s *Sender) Created(c echo.Context, statusCode int, message string, data interface{}) error {
	log.Printf("[INFO] CODE:%d, MESSAGE:%s", statusCode, message)
	return c.JSON(statusCode, Response{
		Code:    statusCode,
		Success: true,
		Message: message,
		Data:    &data,
	})
}
func (s *Sender) Error(c echo.Context, statusCode int, message string, errCode string) error {
	log.Printf("[ERROR] CODE:%d, MESSAGE:%s, ERROR:%s", statusCode, message, errCode)
	return c.JSON(statusCode, Response{
		Code:    statusCode,
		Success: false,
		Message: message,
		Data:    nil,
		Error:   errCode,
	})
}
func (s *Sender) SuccessWithToken(c echo.Context, statusCode int, message string, data interface{}, token string) error {
	log.Printf("[INFO] CODE:%d, MESSAGE:%s", statusCode, message)
	c.Response().Header().Add("Authorization", "Bearer "+token)
	return c.JSON(statusCode, Response{
		Code:    statusCode,
		Success: true,
		Message: message,
		Data:    &data,
	})
}
