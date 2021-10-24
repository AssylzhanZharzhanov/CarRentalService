package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advert struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Status           string             `json:"status" bson:"status"`
	City             string             `json:"city" bson:"city"`
	Category         string             `json:"category" bson:"category" form:"category"`
	Phone            string             `json:"phone" bson:"phone"`
	Title            string             `json:"title" bson:"title" form:"title"`
	TitleSearch      []string           `json:"title_search" bson:"title_search"`
	Description      string             `json:"description" bson:"description" form:"description"`
	RentType         string             `json:"rent_type" bson:"rent_type" form:"rent_type"`
	Price            int                `json:"price" bson:"price" form:"price"`
	Images           []Image `json:"images" bson:"images" form:"images"`
	HasAdvertisement bool               `json:"has_advertisement" bson:"has_advertisement,omitempty"`
	Advertisement    Advertisement      `json:"advertisement" bson:"advertisement,omitempty"`
	Feedbacks        []Feedback         `json:"feedbacks" bson:"feedbacks,omitempty"`
	TotalRating      float64            `json:"total_rating" bson:"total_rating"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}

type AdvertOutput struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status           string             `json:"status" bson:"status"`
	City             string             `json:"city" bson:"city"`
	Category         string             `json:"category" bson:"category"`
	Phone            string             `json:"phone" bson:"phone"`
	Title            string             `json:"title" bson:"title" form:"title"`
	Description      string             `json:"description" bson:"description"`
	RentType         string             `json:"rent_type" bson:"rent_type"`
	Price            int                `json:"price" bson:"price"`
	Images           []Image 			`json:"images" bson:"images"`
	//HasAdvertisement bool               `json:"has_advertisement" bson:"has_advertisement"`
	//Advertisement    Advertisement      `json:"advertisement" bson:"advertisement,omitempty"`
	Feedbacks        []Feedback         `json:"feedbacks" bson:"feedbacks"`
	TotalRating      float64            `json:"total_rating" bson:"total_rating"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
}

type AdvertInput struct {
	Title            string             ` bson:"title" form:"title" binding:"required"`
	Status           string             `json:"status" bson:"status"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	City             string             ` bson:"city" form:"city" binding:"required"`
	Category         string             ` bson:"category" form:"category" binding:"required"`
	Phone            string             ` bson:"phone" form:"phone" binding:"required"`
	Description      string             ` bson:"description" form:"description" binding:"required"`
	RentType         string             ` bson:"rent_type" form:"rent_type" binding:"required"`
	Price            int                ` bson:"price" form:"price" binding:"required"`
	Images            []Image 			` bson:"images" form:"images"`
	Feedbacks        []Feedback         `json:"feedbacks" bson:"feedbacks,omitempty"`
	HasAdvertisement bool               `json:"has_advertisement" bson:"has_advertisement,omitempty"`
	TitleSearch      []string           `json:"title_search" bson:"title_search"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}

func ToAdvert(advert *Advert) *Advert {
	return &Advert{
		ID:               advert.ID,
		Category:         advert.Category,
		Title:            advert.Title,
		Description:      advert.Description,
		RentType:         advert.RentType,
		Price:            advert.Price,
		Images:           advert.Images,
		HasAdvertisement: advert.HasAdvertisement,
		Feedbacks: 		  advert.Feedbacks,
		TotalRating:      advert.TotalRating,
		CreatedAt:        advert.CreatedAt,
	}
}

type UpdateAdvertInput struct {
	Status           string        `json:"status,omitempty" bson:"status,omitempty"`
	Category         string        `json:"category,omitempty" bson:"category,omitempty"`
	Title            string        `json:"title,omitempty" bson:"title,omitempty"`
	Description      string        `json:"description,omitempty" bson:"description,omitempty"`
	RentType         string        `json:"rent_type,omitempty" bson:"rent_type,omitempty"`
	Price            int32         `json:"price,omitempty" bson:"price,omitempty"`
	HasAdvertisement bool          `json:"has_advertisement,omitempty" bson:"has_advertisement,omitempty"`
	Advertisement    Advertisement `json:"advertisement,omitempty" bson:"advertisement,omitempty"`
	Feedbacks        []Feedback    `json:"feedbacks,omitempty" bson:"feedbacks,omitempty"`
	Reason           string        `json:"reason,omitempty" bson:"reason,omitempty"`
}
