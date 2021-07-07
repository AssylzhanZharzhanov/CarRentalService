package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Advert struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Category string `json:"category" bson:"category"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	RentType string `json:"rent_type" bson:"rent_type"`
	Price int32 `json:"price" bson:"price"`
	Images []string `json:"images" bson:"images"`
	HasAdvertisement bool `json:"has_advertisement"`
	AdvertisementType string `json:"advertisement_type" bson:"advertisement_type"`
	Feedbacks []Feedback `json:"feedbacks" bson:"feedbacks"`
}