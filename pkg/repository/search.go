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

func (r *SearchMongo) GetAdverts(ctx context.Context, name string) ([]models.AdvertOutput, error) {
	adverts := make([]models.AdvertOutput, 0)
	searchValue := fmt.Sprintf("^%s", strings.ToLower(name))

	cur, err := r.db.Collection(advertsCollection).Find(ctx, bson.M{"title_search": bson.M{"$regex": searchValue}})

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
	searchValue := fmt.Sprintf("^%s", value)

	cur, err := r.db.Collection(carBrandsCollection).Find(ctx, bson.M{"text_search": bson.M{"$regex": searchValue}})
	//searchStage := bson.D{{"$search", bson.D{{"autocomplete", bson.D{{"path", "brand"}, {"query", brand}}}}}}
	//limitStage := bson.D{{"$limit", 10}}
	//projectStage := bson.D{{"$project", bson.D{{"_id", 0}, {"brand", 1}}}}

	//showInfoCursor, err := r.db.Aggregate(ctx, bson.A{
	//	bson.M{"$search": bson.M{"autocomplete": bson.M{
	//		"query": brand,
	//		"path": "full_name",
	//		"tokenOrder": "any",
	//	}}},
	//	//bson.M{
	//	//	"$limit": 10,
	//	//},
	//})
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
