package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	db *mongo.Database
}

func (c *CategoryRepository) AddCategory(ctx context.Context, category models.Category) error {
	panic("implement me")
}

func (c *CategoryRepository) GetCategories(ctx context.Context) (models.Category, error) {
	panic("implement me")
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	panic("implement me")
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{db:db}
}
