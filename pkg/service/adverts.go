package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type AdvertService struct {
	repo repository.Adverts
}

func NewAdvertService(repo *repository.Repository) *AdvertService {
	return &AdvertService{repo: repo.Adverts}
}

func (s *AdvertService) CreateAdvert(ctx context.Context, advert models.Advert) (string, error) {
	return s.repo.CreateAdvert(ctx, advert)
}

func (s *AdvertService) GetAllAdverts(ctx context.Context) error {
	return s.repo.GetAllAdverts(ctx)
}

func GetAdvertById(ctx context.Context, id string)  {

}