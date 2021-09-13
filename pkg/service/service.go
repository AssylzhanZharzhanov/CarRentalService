package service

import (
	"context"

	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Users interface {
}

type Authentication interface {
	SignUp(ctx context.Context, user models.User) (string, error)
	SignIn(ctx context.Context, user models.User) (string, error)
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.Advert) (string, error)
	GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error)
	GetAdvertById(ctx context.Context, id string) (*models.Advert, error)
	UpdateAdvert(ctx context.Context, id string, advert models.UpdateAdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
}

type Search interface {
	GetCarModels(ctx context.Context) error
}

type Images interface {
	UploadImage(urls []string) error
	GetImageById(ctx context.Context, id string) error
}

type Feedback interface {
	AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error
}

type Category interface {
	AddCategory(ctx context.Context, category models.Category) error
	GetCategories(ctx context.Context) error
	DeleteCategory(ctx context.Context, id string) error
}

type Service struct {
	Authentication
	Adverts
	Users
	Images
	Search
	Feedback
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos),
		Adverts:        NewAdvertService(repos),
		Images:         NewImageService(repos),
		Search:         NewSearchService(repos),
		Feedback: 		NewFeedbackService(repos),
	}
}
