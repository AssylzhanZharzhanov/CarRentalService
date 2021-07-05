package service

import "gitlab.com/zharzhanov/region/pkg/repository"

type Authentication interface {

}

type Adverts interface {

}

type Users interface {

}

type Service struct {
	Authentication
	Adverts
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}