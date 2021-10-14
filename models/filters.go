package models

type Category struct {
	Name string             `json:"name" bson:"name"`
}

type City struct {
	Name string             `json:"name" bson:"name"`
}

type RentTypes struct {
	Name string             `json:"name" bson:"name"`
}

type Price struct {
	Name string				`json:"name" bson:"name"`
	Min  int  				`json:"min,omitempty" bson:"min"`
	Max  int 				`json:"max,omitempty" bson:"max"`
}
