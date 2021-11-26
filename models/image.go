package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID       primitive.ObjectID `json:"id" bson:"_id" bson:"id"`
	Url      string             `json:"url" bson:"url"`
}
