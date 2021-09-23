package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type SearchMongo struct {
	db *mongo.Collection
}

func NewSearchMongo(db *mongo.Database, collection string) *SearchMongo {
	return &SearchMongo{db: db.Collection(collection)}
}

func (r *SearchMongo) GetCarModels(ctx context.Context) error {
	return nil
}

func (r *SearchMongo) SpellChecker() error {
	return nil
}
