package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advert struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Status           string             `json:"status" bson:"status"`
	Mark			 string             `json:"mark" bson:"mark"`
	Model			 string             `json:"model" bson:"model"`
	Category         string             `json:"category" bson:"category"`
	Class         	 string             `json:"class" bson:"class"`
	BodyStyle        string             `json:"body_style" bson:"body_style"`
	Transmission     string             `json:"transmission" bson:"transmission"`
	Year             string             `json:"year" bson:"year"`
	City             string             `json:"city" bson:"city"`
	Title            string             `json:"title" bson:"title" form:"title"`
	TitleSearch      []string           `json:"title_search" bson:"title_search"`
	Description      string             `json:"description" bson:"description"`
	RentType         string             `json:"rent_type" bson:"rent_type"`
	Price            int                `json:"price" bson:"price"`
	Phone            string             `json:"phone" bson:"phone"`
	Images           []Image 		    `json:"images" bson:"images"`
	HasAdvertisement bool               `json:"has_advertisement" bson:"has_advertisement"`
	//Advertisement    []Advertisement    `json:"advertisement" bson:"advertisement"`
	Feedbacks        []Feedback         `json:"feedbacks" bson:"feedbacks"`
	TotalRating      float64            `json:"total_rating" bson:"total_rating"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt "`
}

type AdvertInput struct {
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
	Title            string             `json:"title,omitempty" bson:"title,omitempty" form:"title"`
	Mark			 string             `json:"mark" bson:"mark" form:"mark"`
	Model			 string             `json:"model" bson:"model" form:"model"`
	Class         	 string             `json:"class" bson:"class"`
	BodyStyle        string             `json:"body_style" bson:"body_style"`
	Transmission     string             `json:"transmission" bson:"transmission"`
	Year             string             `json:"year" bson:"year"`
	City             string             `json:"city" bson:"city"`
	Category         string             `json:"category,omitempty" bson:"category,omitempty" form:"category" `
	Description      string             `json:"description,omitempty" bson:"description,omitempty" form:"description"`
	RentType         string             `json:"rent_type,omitempty" bson:"rent_type,omitempty" form:"rent_type"`
	Price            int                `json:"price,omitempty" bson:"price,omitempty" form:"price" `
	Phone            string             `json:"phone,omitempty" bson:"phone" form:"phone" `
	Images           []Image 			`json:"images,omitempty" bson:"images,omitempty" form:"images"`
	Feedbacks        []Feedback         `json:"feedbacks,omitempty" bson:"feedbacks,omitempty"`
	HasAdvertisement bool               `json:"has_advertisement,omitempty" bson:"has_advertisement,omitempty"`
	TitleSearch      []string           `json:"title_search,omitempty" bson:"title_search,omitempty"`
	CreatedAt        time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
