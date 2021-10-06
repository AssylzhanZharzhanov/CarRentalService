package models

type User struct {
	Id         int32  `json:"id" bson:"_id"`
	Phone      string `json:"phone" bson:"phone"`
}

type UserInput struct {
	FirstName  string `json:"firstName" bson:"first_name"`
	SecondName string `json:"secondName" bson:"second_name"`
	Username   string `json:"username" bson:"username"`
	Password   string `json:"password" bson:"password"`
	Avatar	   string `json:"avatar" bson:"avatar"`
}

