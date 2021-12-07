package repository

import (
	"context"
	"fmt"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const carBrandsCollection = "car_brands"


type SearchMongo struct {
	db *mongo.Database
}

func NewSearchMongo(db *mongo.Database) *SearchMongo {
	return &SearchMongo{db: db}
}

func (r *SearchMongo) GetAdverts(ctx context.Context, name string) ([]models.Advert, error) {
	adverts := make([]models.Advert, 0)
	searchValue := fmt.Sprintf("\"%s\"", strings.ToLower(name))

	cur, err := r.db.Collection(advertsCollection).Find(ctx, bson.M{"$text": bson.M{"$search": searchValue, "$caseSensitive": false}})
	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, nil
}

func (r *SearchMongo) GetCarModels(ctx context.Context, value string) ([]models.CarModels, error) {
	carModels := make([]models.CarModels, 0)
	searchValue := fmt.Sprintf("\"%s\"", value)

	cur, err := r.db.Collection(carBrandsCollection).Find(ctx, bson.M{"$text": bson.M{"$search": searchValue, "$caseSensitive": false}})

	if err != nil {
		return carModels, err
	}
	if err = cur.All(ctx, &carModels); err != nil {
		return carModels, err
	}

	return carModels, nil
}

func (r *SearchMongo) SpellChecker() error {
	return nil
}
