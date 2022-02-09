package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdvertisementsMongo struct {
	db *mongo.Database
}

func (r *AdvertisementsMongo) CreateAdvertisement(ctx context.Context, advertisement models.AdvertisementInput) error {
	panic("implement me")
}

func (r *AdvertisementsMongo) GetAdvertisements(ctx context.Context) ([]models.Advertisement, error) {
	advertisements := make([]models.Advertisement, 0)
	filter := bson.M{}

	cur, err := r.db.Collection(advertisementsCollection).Find(ctx, filter)
	if err != nil {
		return advertisements, err
	}

	if err = cur.All(ctx, &advertisements); err != nil {
		return advertisements, err
	}

	return advertisements, err
}

func (r *AdvertisementsMongo) GetAdvertisementByID(ctx context.Context) (models.Advertisement, error) {
	panic("implement me")
}

func (r *AdvertisementsMongo) UpdateAdvertisement(ctx context.Context, id string, advertisement models.AdvertisementInput) error {
	panic("implement me")
}

func (r *AdvertisementsMongo) DeleteAdvertisement(ctx context.Context, id string) error {
	panic("implement me")
}

func NewAdvertisementMongo(db *mongo.Database) *AdvertisementsMongo {
	return &AdvertisementsMongo{db: db}
}
