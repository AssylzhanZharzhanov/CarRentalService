package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type FilterService struct {
	repo repository.Filters
}

func (s *FilterService) AddCategory(ctx context.Context, category models.Category) error {
	return s.repo.AddCategory(ctx, category)
}

func (s *FilterService) GetCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *FilterService) DeleteCategory(ctx context.Context, id string) error {
	return s.repo.DeleteCategory(ctx, id)
}

func NewFilterService(repository *repository.Repository) *FilterService {
	return &FilterService{
		repo: repository.Filters,
	}
}


