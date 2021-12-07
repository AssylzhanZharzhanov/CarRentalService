package service

import (
	"context"

	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Users interface {
}

type Admin interface {
	GetUser(ctx context.Context, phone string) (string, error)
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.AdvertInput, imageUrl []string, userId string) (string, error)
	GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error)
	GetAdvertById(ctx context.Context, id string) (models.Advert, error)
	GetMyAdverts(ctx context.Context, userId string) ([]models.Advert, error)
	UpdateAdvert(ctx context.Context, id string, advert models.AdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
}

type Advertisements interface {
	CreateAdvertisement(ctx context.Context, advertisement models.AdvertisementInput) error
	GetAdvertisements(ctx context.Context) ([]models.Advertisement, error)
	GetAdvertisementByID(ctx context.Context) (models.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, id string, advertisement models.AdvertisementInput) error
	DeleteAdvertisement(ctx context.Context, id string) error
}

type Bookmarks interface {
	AddUserBookmark(ctx context.Context, userId string, advertId string) error
	GetUserBookmarks(ctx context.Context, userId string) ([]models.Advert, error)
	RemoveUserBookmark(ctx context.Context, userId string, advertId string) error
}

type Authentication interface {
	VerifyCode(ctx context.Context, code string) (string, error)
	SendSMS(ctx context.Context, phone string) (string, error)
	SignUp(ctx context.Context, user models.User) (string, error)
	SignIn(ctx context.Context, user models.User) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Search interface {
	GetCarModels(ctx context.Context, brand string) ([]models.CarModels, error)
	GetAdverts(ctx context.Context, name string) ([]models.Advert, error)
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
	UpdateFeedback(ctx context.Context, feedbackId string, feedback models.Feedback) error
	DeleteFeedback(ctx context.Context, feedbackId string) error
}

type Filters interface {
	AddCategory(ctx context.Context, category models.Category) error
	GetCategories(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, name string) error

	AddCity(ctx context.Context, city models.City) error
	GetCities(ctx context.Context) ([]models.City, error)
	DeleteCity(ctx context.Context, name string) error

	AddRentType(ctx context.Context, city models.RentTypes) error
	GetRentTypes(ctx context.Context) ([]models.RentTypes, error)
	DeleteRentType(ctx context.Context, name string) error

	AddPrice(ctx context.Context, city models.Price) error
	GetPrices(ctx context.Context) ([]models.Price, error)
	DeletePrices(ctx context.Context, name string) error

	AddStatus(ctx context.Context, status models.Status) error
	GetStatus(ctx context.Context) ([]models.Status, error)
	DeleteStatus(ctx context.Context, name string) error
}

type Service struct {
	Authentication
	Adverts
	Advertisements
	Images
	Search
	Feedback
	Filters
	Bookmarks
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos),
		Adverts:        NewAdvertService(repos),
		Advertisements: NewAdvertisementService(repos),
		Images:         NewImageService(repos),
		Search:         NewSearchService(repos),
		Feedback: 		NewFeedbackService(repos),
		Filters:        NewFilterService(repos),
		Bookmarks:      NewBookmarkService(repos),
		Admin:          NewAdminService(repos),
	}
}
