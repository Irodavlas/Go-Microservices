package model

// defines structs and db connections
type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
