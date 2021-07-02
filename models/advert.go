package models

type Advert struct {
	ID int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}