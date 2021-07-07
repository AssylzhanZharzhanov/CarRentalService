package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdvertMongo struct {
	db *mongo.Database
}

func NewAdvertMongo(db *mongo.Database) *AdvertMongo {
	return &AdvertMongo{db:db}
}

func (r *AdvertMongo) CreateAdvert(ctx context.Context, advert models.Advert) (string, error) {
	return "", nil
}

func (r *AdvertMongo) GetAllAdverts(ctx context.Context) error {
	return nil
}