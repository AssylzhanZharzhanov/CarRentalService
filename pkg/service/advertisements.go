package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type AdvertisementService struct {
	repo repository.Advertisements
}

func (s *AdvertisementService) CreateAdvertisement(ctx context.Context, advertisement models.AdvertisementInput) error {
	panic("implement me")
}

func (s *AdvertisementService) GetAdvertisements(ctx context.Context) ([]models.Advertisement, error) {
	return s.repo.GetAdvertisements(ctx)
}

func (s *AdvertisementService) GetAdvertisementByID(ctx context.Context) (models.Advertisement, error) {
	panic("implement me")
}

func (s *AdvertisementService) UpdateAdvertisement(ctx context.Context, id string, advertisement models.AdvertisementInput) error {
	panic("implement me")
}

func (s *AdvertisementService) DeleteAdvertisement(ctx context.Context, id string) error {
	panic("implement me")
}

func NewAdvertisementService(repository *repository.Repository) *AdvertisementService {
	return &AdvertisementService{repo: repository.Advertisements}
}
