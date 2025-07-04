package model

type ResponseGetUser struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    User   `json:"data,omitempty"`
}
