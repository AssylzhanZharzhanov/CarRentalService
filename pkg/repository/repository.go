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

type Authentication interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, user models.User) (string, error)
}

type Feedback interface {
	AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error
	UpdateRating(ctx context.Context, advertId string) error
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.Advert) (string, error)
	GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error)
	GetAdvertById(ctx context.Context, id string) (*models.Advert, error)
	UpdateAdvert(ctx context.Context, id string, advert models.UpdateAdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
}

type Search interface {
	SpellChecker() error
	GetCarModels(ctx context.Context) error
}

type Users interface {
}

type Images interface {
	UploadImage(ctx context.Context, advertId string, url string) error
	DeleteImage(ctx context.Context, imageId string, advertId string) error
	GetImageById(ctx context.Context, id string) (models.Image, error)
}

type Category interface {
	AddCategory(ctx context.Context, category models.Category) error
	GetCategories(ctx context.Context) (models.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type Repository struct {
	Authentication
	Adverts
	Users
	Images
	Search
	Feedback
	Category
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
		Adverts:        NewAdvertMongo(db, advertsCollection),
		Images:         NewImageMongo(db),
		Search:         NewSearchMongo(db, advertsCollection),
		Feedback:       NewFeedbackMongo(db, advertsCollection),
		Category:       NewCategoryRepository(db),
	}
}
