package models

type User struct {
	id int32 `json:"id"`
	Name string `json:"username"`
	Username string `json:"username"`
	Password string `json:"password"`
}