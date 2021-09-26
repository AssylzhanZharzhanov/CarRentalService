package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilterRepository struct {
	db *mongo.Database
}

func (c *FilterRepository) AddCategory(ctx context.Context, category models.Category) error {
	panic("implement me")
}

func (c *FilterRepository) GetCategories(ctx context.Context) (models.Category, error) {
	panic("implement me")
}

func (c *FilterRepository) DeleteCategory(ctx context.Context, id string) error {
	panic("implement me")
}

func NewFilterRepository(db *mongo.Database) *FilterRepository {
	return &FilterRepository{db:db}
}
