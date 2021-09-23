package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bookmark struct {
	ID primitive.ObjectID
	UserID primitive.ObjectID
	Adverts []primitive.ObjectID
}

type BookMarkAdvert struct {
	AdvertID primitive.ObjectID
	Title string
	Region string
	Price string
}