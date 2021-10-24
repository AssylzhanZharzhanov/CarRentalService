package models

type CarModels struct {
	Mark        string   `json:"mark,omitempty" bson:"brand,omitempty"`
	Model       string   `json:"model,omitempty" bson:"model,omitempty"`
	FullName    string   `json:"full_name" bson:"full_name"`
	//Count       int      `json:"count" bson:"count"`
}
