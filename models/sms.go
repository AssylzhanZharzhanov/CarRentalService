package models

import "time"

type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type Code struct {
	Phone     string            `json:"phone" bson:"phone"`
	Code      string            `json:"code" bson:"code"`
	ExpiresAt time.Time         `json:"expires_at" bson:"expires_at"`
}

type InputCode struct {
	Code  string            `json:"code" bson:"code"`
}

type GeneratedCode struct {
	Code string `json:"code"`
}

