package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strings"
	"time"
)

const (
	defaultStatusValue = "На модерации"
)

type AdvertService struct {
	repo repository.Adverts
}

func NewAdvertService(repo *repository.Repository) *AdvertService {
	return &AdvertService{repo: repo.Adverts}
}

func (s *AdvertService) GetMyAdverts(ctx context.Context, userId string) ([]models.Advert, error) {
	return s.repo.GetMyAdverts(ctx, userId)
}

func (s *AdvertService) CreateAdvert(ctx context.Context, advert models.AdvertInput, imageUrl []string, userId string) (string, error) {

	userObjId, _ := primitive.ObjectIDFromHex(userId)

	advert.UserID = userObjId
	advert.CreatedAt = time.Now()
	advert.HasAdvertisement = false
	advert.TitleSearch = strings.Fields(strings.ToLower(advert.Title))
	advert.Feedbacks = make([]models.Feedback, 0)
	advert.Images = make([]models.Image, 0)
	advert.Status = defaultStatusValue

	advertId, err := s.repo.CreateAdvert(ctx, advert)
	if err != nil {
		return "", err
	}

	err = s.repo.UploadImage(ctx, advertId, imageUrl)
	if err != nil {
		log.Println(err.Error())
	}

	return advertId, nil
}

func (s *AdvertService) GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error) {
	return s.repo.GetAllAdverts(ctx, filter)
}

func (s *AdvertService) GetAdvertById(ctx context.Context, id string) (models.Advert, error) {
	return s.repo.GetAdvertById(ctx, id)
}

func (s *AdvertService) UpdateAdvert(ctx context.Context, id string, advert models.AdvertInput) error {
	return s.repo.UpdateAdvert(ctx, id, advert)
}

func (s *AdvertService) DeleteAdvert(ctx context.Context, advertId string) error {
	return s.repo.DeleteAdvert(ctx, advertId)
}
