package service

import (
	"context"

	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Users interface {
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.AdvertInput, imageUrl []string, userId string) (string, error)
	GetAllAdverts(ctx context.Context, filter bson.M) ([]models.AdvertOutput, error)
	GetAdvertById(ctx context.Context, id string) (models.AdvertOutput, error)
	UpdateAdvert(ctx context.Context, id string, advert models.UpdateAdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
}

type Bookmarks interface {
	AddUserBookmark(ctx context.Context, userId string, advertId string) error
	GetUserBookmarks(ctx context.Context, userId string) ([]models.Advert, error)
}

type Authentication interface {
	VerifyCode(ctx context.Context, code string) (string, error)
	SendSMS(ctx context.Context, phone string) (string, error)
	SignUp(ctx context.Context, user models.User) (string, error)
	SignIn(ctx context.Context, user models.User) (string, error)
	ParseToken(accessToken string) (string, error)
}


type Search interface {
	GetCarModels(ctx context.Context) error
}

type Images interface {
	UploadImage(ctx context.Context, advertId string, url string) error
	UploadMultipleImages(ctx context.Context, urls []string) error
	GetImageById(ctx context.Context, id string) (models.Image, error)
	DeleteImage(ctx context.Context, imageId string, advertId string) error
}

type Feedback interface {
	AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error
	GetFeedbackByUserId(ctx context.Context, feedbackId string) (*models.Feedback, error)
}

type Filters interface {
	AddCategory(ctx context.Context, category models.Category) error
	GetCategories(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type Service struct {
	Authentication
	Adverts
	Images
	Search
	Feedback
	Filters
	Bookmarks
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos),
		Adverts:        NewAdvertService(repos),
		Images:         NewImageService(repos),
		Search:         NewSearchService(repos),
		Feedback: 		NewFeedbackService(repos),
		Filters:        NewFilterService(repos),
		Bookmarks:      NewBookmarkService(repos),
	}
}
