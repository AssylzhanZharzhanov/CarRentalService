package service

import (
	"context"

	"gitlab.com/zharzhanov/region/pkg/repository"
)

type SearchService struct {
	repo repository.Search
}

func NewSearchService(repository *repository.Repository) *SearchService {
	return &SearchService{repo: repository.Search}
}

func (s *SearchService) GetCarModels(ctx context.Context) error {

	return s.repo.GetCarModels(ctx)
}
