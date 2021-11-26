package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advert struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Status           string             `json:"status" bson:"status"`
	Reason           string       		`json:"reason,omitempty" bson:"reason,omitempty"`
	City             string             `json:"city" bson:"city"`
	Category         string             `json:"category" bson:"category" form:"category"`
	Phone            string             `json:"phone" bson:"phone"`
	Title            string             `json:"title" bson:"title" form:"title"`
	TitleSearch      []string           `json:"title_search" bson:"title_search"`
	Description      string             `json:"description" bson:"description" form:"description"`
	RentType         string             `json:"rent_type" bson:"rent_type" form:"rent_type"`
	Price            int                `json:"price" bson:"price" form:"price"`
	Images           []Image 		    `json:"images" bson:"images" form:"images"`
	HasAdvertisement bool               `json:"has_advertisement" bson:"has_advertisement,omitempty"`
	Advertisement    Advertisement      `json:"advertisement" bson:"advertisement,omitempty"`
	Feedbacks        []Feedback         `json:"feedbacks" bson:"feedbacks,omitempty"`
	TotalRating      float64            `json:"total_rating" bson:"total_rating"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}

type AdvertInput struct {
	Title            string             `json:"title,omitempty" bson:"title,omitempty" form:"title" binding:"required"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	City             string             `json:"city,omitempty" bson:"city,omitempty" form:"city" `
	Category         string             `json:"category,omitempty" bson:"category,omitempty" form:"category" `
	Phone            string             `json:"phone,omitempty" bson:"phone" form:"phone,omitempty" `
	Description      string             `json:"description,omitempty" bson:"description,omitempty" form:"description"`
	RentType         string             `json:"rent_type,omitempty" bson:"rent_type,omitempty" form:"rent_type"`
	Price            int                `json:"price,omitempty" bson:"price,omitempty" form:"price" `
	Images           []Image 			`json:"images,omitempty" bson:"images,omitempty" form:"images"`
	Feedbacks        []Feedback         `json:"feedbacks,omitempty" bson:"feedbacks,omitempty"`
	HasAdvertisement bool               `json:"has_advertisement,omitempty" bson:"has_advertisement,omitempty"`
	TitleSearch      []string           `json:"title_search,omitempty" bson:"title_search,omitempty"`
	CreatedAt        time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
