package models

type CarModels struct {
	Model       string   `json:"model" bson:"model"`
	ModelSearch []string `json:"modelSearch" bson:"modelSearch"`
	Count       int      `json:"count" bson:"count"`
}
