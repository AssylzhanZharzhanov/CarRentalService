package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advert struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Category          string             `json:"category" bson:"category" form:"category"`
	Title             string             `json:"title" bson:"title" form:"title"`
	Description       string             `json:"description" bson:"description" form:"description"`
	RentType          string             `json:"rent_type" bson:"rent_type" form:"rent_type"`
	Price             int32              `json:"price" bson:"price" form:"price"`
	Images            []string           `json:"images" bson:"images" form:"images"`
	HasAdvertisement  bool               `json:"has_advertisement" bson:"has_advertisement,omitempty"`
	Advertisement 	  Advertisement 	 `json:"advertisement" bson:"advertisement,omitempty"`
	Feedbacks         []Feedback         `json:"feedbacks,omitempty" bson:"feedbacks,omitempty"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}

func ToAdvert(advert *Advert) *Advert {
	return &Advert{
		ID: advert.ID,
		Category: advert.Category,
		Title: advert.Title,
		Description: advert.Description,
		RentType: advert.RentType,
		Price: advert.Price,
		Images: advert.Images,
		HasAdvertisement: advert.HasAdvertisement,
		CreatedAt: advert.CreatedAt,
	}
}

type UpdateAdvertInput struct {
	Category          string             `json:"category" bson:"category"`
	Title             string             `json:"title" bson:"title"`
	Description       string             `json:"description" bson:"description"`
	RentType          string             `json:"rent_type" bson:"rent_type"`
	Price             int32              `json:"price" bson:"price"`
	Images            []string           `json:"images" bson:"images"`
	HasAdvertisement  bool               `json:"has_advertisement" bson:"has_advertisement,omitempty"`
	Advertisement 	  Advertisement 	 `json:"advertisement" bson:"advertisement,omitempty"`
	Feedbacks         []Feedback         `json:"feedbacks,omitempty" bson:"feedbacks,omitempty"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}
