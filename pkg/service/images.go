package service

import (
	"context"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type ImageService struct {
	repo repository.Images
}

func NewImageService(repository *repository.Repository) *ImageService {
	return &ImageService{
		repo: repository.Images,
	}
}

func (s *ImageService) UploadImage(urls []string) error {
	return s.repo.UploadImage(urls)
}
func (s *ImageService) GetImageById(ctx context.Context, id string) error {
	return s.repo.GetImageById(ctx, id)
}

