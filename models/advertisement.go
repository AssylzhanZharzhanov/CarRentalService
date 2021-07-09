package models

import "time"

type Advertisement struct {
	Type            string    `json:"type" bson:"type"`
	Status          string    `json:"status" bson:"status"`
	InspirationDate time.Time `json:"inspiration_date" bson:"inspiration_date"`
	ExpirationDate  time.Time `json:"expiration_date" bson:"expiration_date"`
}
