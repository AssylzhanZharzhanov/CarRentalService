package models

type CarModels struct {
	Model       string   `json:"model" bson:"model"`
	Mark        string   `json:"mark" bson:"mark"`
	Count       int      `json:"count" bson:"count"`
}
