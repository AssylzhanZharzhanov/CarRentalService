package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type FilterService struct {

}

func (c *FilterService) AddCategory(ctx context.Context, category models.Category) error {
	panic("implement me")
}

func (c *FilterService) GetCategories(ctx context.Context) (models.Category, error) {
	panic("implement me")
}

func (c *FilterService) DeleteCategory(ctx context.Context, id string) error {
	panic("implement me")
}

func NewFilterService(repository *repository.Repository) *FilterService {
	return &FilterService{}
}


