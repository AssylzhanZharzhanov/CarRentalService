package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type ImageService struct {
	repo repository.Images
}

func (s *ImageService) UploadImage(ctx context.Context, advertId string, url string) error {
	return s.repo.UploadImage(ctx, advertId, url)
}

func (s *ImageService) UploadMultipleImages(ctx context.Context, urls []string) error {
	panic("implement me")
}

func (s *ImageService) DeleteImage(ctx context.Context, imageId string, advertId string) error {
	return s.repo.DeleteImage(ctx, imageId, advertId)
}

func (s *ImageService) GetImageById(ctx context.Context, id string) (models.Image, error) {
	return s.repo.GetImageById(ctx, id)
}

func NewImageService(repository *repository.Repository) *ImageService {
	return &ImageService{
		repo: repository.Images,
	}
}

