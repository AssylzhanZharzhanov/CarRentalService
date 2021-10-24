package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type SearchMongo struct {
	db *mongo.Collection
}

func NewSearchMongo(db *mongo.Database, collection string) *SearchMongo {
	return &SearchMongo{db: db.Collection(collection)}
}

func (r *SearchMongo) GetCarModels(ctx context.Context, brand string) ([]models.CarModels, error) {
	carModels := make([]models.CarModels, 0)

	//searchStage := bson.D{{"$search", bson.D{{"autocomplete", bson.D{{"path", "brand"}, {"query", brand}}}}}}
	limitStage := bson.D{{"$limit", 10}}
	projectStage := bson.D{{"$project", bson.D{{"_id", 0}, {"brand", 1}}}}

	showInfoCursor, err := r.db.Aggregate(ctx, mongo.Pipeline{limitStage, projectStage})
	if err != nil {
		return carModels, err
	}

	if err = showInfoCursor.All(ctx, &carModels); err != nil {
		return carModels, err
	}

	return carModels, nil
}

func (r *SearchMongo) SpellChecker() error {
	return nil
}
