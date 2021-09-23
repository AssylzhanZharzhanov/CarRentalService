package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Advertisement struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Type            string    `json:"type" bson:"type"`
	Status          string    `json:"status" bson:"status"`
	InspirationDate time.Time `json:"inspiration_date" bson:"inspiration_date"`
	ExpirationDate  time.Time `json:"expiration_date" bson:"expiration_date"`
}
