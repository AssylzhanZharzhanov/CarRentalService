package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type CategoryService struct {

}

func (c *CategoryService) AddCategory(ctx context.Context, category models.Category) error {
	panic("implement me")
}

func (c *CategoryService) GetCategories(ctx context.Context) (models.Category, error) {
	panic("implement me")
}

func (c *CategoryService) DeleteCategory(ctx context.Context, id string) error {
	panic("implement me")
}

func NewCategoryService(repository *repository.Repository) *CategoryService {
	return &CategoryService{}
}


