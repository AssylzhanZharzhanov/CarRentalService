package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feedback struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Comment string `json:"comment" bson:"comment"`
	Rating int `json:"rating" bson:"rating"`
}