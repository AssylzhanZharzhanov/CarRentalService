package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	categoriesCollection = "categories"
)

type FilterRepository struct {
	db *mongo.Database
}

func (r *FilterRepository) AddCategory(ctx context.Context, category models.Category) error {
	_, err := r.db.Collection(categoriesCollection).InsertOne(ctx, category)
	return err
}

func (r *FilterRepository) GetCategories(ctx context.Context) ([]models.Category, error) {
	categories := make([]models.Category, 0)

	filter := bson.M{}
	cur, err := r.db.Collection(categoriesCollection).Find(ctx, filter)

	if err = cur.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *FilterRepository) DeleteCategory(ctx context.Context, id string) error {
	panic("implement me")
}

func NewFilterRepository(db *mongo.Database) *FilterRepository {
	return &FilterRepository{db:db}
}
