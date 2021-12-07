package service

import "gitlab.com/zharzhanov/region/pkg/repository"

type AdvertisementService struct {

}

func NewAdvertisementService(repo *repository.Repository) *Service {
	return &Service{}
}
