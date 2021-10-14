package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	//ID       int32      			`json:"id" bson:"_id"`
	Phone      string 	  			`json:"phone" bson:"phone"`
	Bookmarks  []primitive.ObjectID `json:"bookmarks" bson:"bookmarks"`
}

type UserInput struct {
	FirstName  string `json:"firstName" bson:"first_name"`
	SecondName string `json:"secondName" bson:"second_name"`
	Username   string `json:"username" bson:"username"`
	Password   string `json:"password" bson:"password"`
	Avatar	   string `json:"avatar" bson:"avatar"`
}

