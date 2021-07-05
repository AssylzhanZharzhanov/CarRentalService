package models

type User struct {
	id int32 `json:"id"`
	FirstName string `json:"firstName"`
	SecondName string `json:"secondName"`
	Username string `json:"username"`
	Password string `json:"password"`
}