package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImageMongo struct {
	db *mongo.Collection
}

func (i *ImageMongo) UploadImage(urls []string) error {
	panic("implement me")
}

func (i *ImageMongo) GetImageById(ctx context.Context, id string) error {
	panic("implement me")
}

func NewImageMongo(db *mongo.Database, collection string) *ImageMongo {
	return &ImageMongo{db: db.Collection(collection)}
}
