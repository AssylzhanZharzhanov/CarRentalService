package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	Id 		  primitive.ObjectID `json:"id" bson:"_id"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Comment   string             `json:"comment" bson:"comment"`
	Rating    int                `json:"rating" bson:"rating"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type FeedbackInput struct {
	Comment   string             `json:"comment,omitempty" bson:"comment"`
	Rating    int                `json:"rating" bson:"rating"`
}

