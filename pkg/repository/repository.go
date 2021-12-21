package repository

import (
	"context"

	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	advertsCollection = "adverts"
	imageCollection = "images"
)

type Admin interface {
}

type Authentication interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, phone string) (string, error)
	CreateCode(ctx context.Context, code models.Code)  error
	VerifyCode(ctx context.Context, code string) (models.Code, error)
}

type Bookmarks interface {
	AddUserBookmark(ctx context.Context, userId string, advertId string) error
	GetUserBookmarks(ctx context.Context, userId string) ([]models.Advert, error)
	RemoveUserBookmark(ctx context.Context, userId string, advertId string) error
}

type Feedback interface {
	AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error
	UpdateRating(ctx context.Context, advertId string) error
	UpdateFeedback(ctx context.Context,  feedbackId string, feedback models.Feedback) error
	DeleteFeedback(ctx context.Context, feedbackId string) error
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.AdvertInput) (string, error)
	GetUserAdverts(ctx context.Context, userId string, status string) ([]models.Advert, error)
	GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error)
	GetAdvertById(ctx context.Context, id string) (models.Advert, error)
	UpdateAdvert(ctx context.Context, id string, advert models.AdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
	UploadImage(ctx context.Context, advertId string, url []string) error
}

type Search interface {
	SpellChecker() error
	GetCarModels(ctx context.Context, brand string) ([]models.CarModels, error)
	GetAdverts(ctx context.Context, name string) ([]models.Advert, error)
}

type Users interface {

}

type Images interface {
	UploadImage(ctx context.Context, advertId string, url string) error
	DeleteImage(ctx context.Context, imageId string, advertId string) error
	GetImageById(ctx context.Context, id string) (models.Image, error)
}

type Filters interface {
	AddCategory(ctx context.Context, category models.Category) error
	GetCategories(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, name string) error

	AddCity(ctx context.Context, city models.City) error
	GetCities(ctx context.Context) ([]models.City, error)
	DeleteCity(ctx context.Context, name string) error

	AddRentType(ctx context.Context, rentType models.RentTypes) error
	GetRentTypes(ctx context.Context) ([]models.RentTypes, error)
	DeleteRentType(ctx context.Context, name string) error

	AddPrice(ctx context.Context, price models.Price) error
	GetPrices(ctx context.Context) ([]models.Price, error)
	DeletePrices(ctx context.Context, name string) error

	AddStatus(ctx context.Context, status models.Status) error
	GetStatuses(ctx context.Context) ([]models.Status, error)
	DeleteStatus(ctx context.Context, name string) error
}

type Repository struct {
	Authentication
	Adverts
	Users
	Images
	Search
	Feedback
	Filters
	Bookmarks
	Admin
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
		Adverts:        NewAdvertMongo(db),
		Images:         NewImageMongo(db),
		Search:         NewSearchMongo(db),
		Feedback:       NewFeedbackMongo(db, advertsCollection),
		Filters:        NewFilterRepository(db),
		Bookmarks:      NewBookmarkMongo(db),
		Admin: 			NewAdminMongo(db),
	}
}
