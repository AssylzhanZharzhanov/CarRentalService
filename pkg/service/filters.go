package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type FilterService struct {
	repo repository.Filters
}

func (s *FilterService) AddStatus(ctx context.Context, status models.Status) error {
	panic("implement me")
}

func (s *FilterService) GetStatus(ctx context.Context) ([]models.Status, error) {
	panic("implement me")
}

func (s *FilterService) DeleteStatus(ctx context.Context, name string) error {
	panic("implement me")
}

func (s *FilterService) AddPrice(ctx context.Context, city models.Price) error {
	return s.repo.AddPrice(ctx, city)
}

func (s *FilterService) GetPrices(ctx context.Context) ([]models.Price, error) {
	return s.repo.GetPrices(ctx)
}

func (s *FilterService) DeletePrices(ctx context.Context, name string) error {
	return s.repo.DeletePrices(ctx, name)
}

func (s *FilterService) AddRentType(ctx context.Context, rentType models.RentTypes) error {
	return s.repo.AddRentType(ctx, rentType)
}

func (s *FilterService) GetRentTypes(ctx context.Context) ([]models.RentTypes, error) {
	return s.repo.GetRentTypes(ctx)
}

func (s *FilterService) DeleteRentType(ctx context.Context, name string) error {
	return s.repo.DeleteRentType(ctx, name)
}

func (s *FilterService) AddCity(ctx context.Context, city models.City) error {
	return s.repo.AddCity(ctx, city)
}

func (s *FilterService) GetCities(ctx context.Context) ([]models.City, error) {
	return s.repo.GetCities(ctx)
}

func (s *FilterService) DeleteCity(ctx context.Context, name string) error {
	return s.repo.DeleteCity(ctx, name)
}

func (s *FilterService) AddCategory(ctx context.Context, category models.Category) error {
	return s.repo.AddCategory(ctx, category)
}

func (s *FilterService) GetCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *FilterService) DeleteCategory(ctx context.Context, name string) error {
	return s.repo.DeleteCategory(ctx, name)
}

func NewFilterService(repository *repository.Repository) *FilterService {
	return &FilterService{
		repo: repository.Filters,
	}
}


