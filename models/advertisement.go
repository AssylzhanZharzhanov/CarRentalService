package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Advertisement struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string    `json:"name" bson:"name"`
	Tag				string    `json:"tag" bson:"tag"`
	Description     string    `json:"description" bson:"description"`
	Status          int       `json:"status" bson:"status"`
	Price           int       `json:"price" bson:"price"`
	IssuedDate      time.Time `json:"issued_date" bson:"inspiration_date"`
	ExpirationDate  time.Time `json:"expiration_date" bson:"expiration_date"`
}

type AdvertisementInput struct {
	Name            string    `json:"name" bson:"name"`
	Description     string    `json:"description" bson:"description"`
	Status          int       `json:"status" bson:"status"`
	Price           int       `json:"price" bson:"price"`
}