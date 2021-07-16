package service

import (
	"context"
	"strings"
	"time"

	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type AdvertService struct {
	repo repository.Adverts
}

func NewAdvertService(repo *repository.Repository) *AdvertService {
	return &AdvertService{repo: repo.Adverts}
}

func (s *AdvertService) CreateAdvert(ctx context.Context, advert models.Advert) (string, error) {
	advert.CreatedAt = time.Now()
	advert.HasAdvertisement = false
	advert.TitleSearch = strings.Fields(strings.ToLower(advert.Title))
	return s.repo.CreateAdvert(ctx, advert)
}

func (s *AdvertService) GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error) {
	return s.repo.GetAllAdverts(ctx, filter)
}

func (s *AdvertService) GetAdvertById(ctx context.Context, id string) (*models.Advert, error) {
	return s.repo.GetAdvertById(ctx, id)
}

func (s *AdvertService) UpdateAdvert(ctx context.Context, id string, advert models.UpdateAdvertInput) error {
	return s.repo.UpdateAdvert(ctx, id, advert)
}

func (s *AdvertService) DeleteAdvert(ctx context.Context, id string) error {
	return s.repo.DeleteAdvert(ctx, id)
}
