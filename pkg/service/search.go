package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"

	"gitlab.com/zharzhanov/region/pkg/repository"
)

type SearchService struct {
	repo repository.Search
}

func NewSearchService(repository *repository.Repository) *SearchService {
	return &SearchService{repo: repository.Search}
}

func (s *SearchService) GetAdverts(ctx context.Context, name string) ([]models.Advert, error) {
	return s.repo.GetAdverts(ctx, name)
}

func (s *SearchService) GetCarModels(ctx context.Context, brand string) ([]models.CarModels, error) {
	return s.repo.GetCarModels(ctx, brand)
}
