package models

type User struct {
	Id         int32  `json:"id" bson:"id"`
	FirstName  string `json:"firstName" bson:"first_name"`
	SecondName string `json:"secondName" bson:"second_name"`
	Username   string `json:"username" bson:"username"`
	Password   string `json:"password" bson:"password"`
	Avatar	   string `json:"avatar" bson:"avatar"`
}
