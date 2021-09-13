package repository

import (
	"context"

	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	advertsCollection = "adverts"
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
	UploadImage(urls []string) error
	GetImageById(ctx context.Context, id string) error
}

type Repository struct {
	Authentication
	Adverts
	Users
	Images
	Search
	Feedback
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
		Adverts:        NewAdvertMongo(db, advertsCollection),
		Images:         NewImageMongo(db, advertsCollection),
		Search:         NewSearchMongo(db, advertsCollection),
		Feedback:       NewFeedbackMongo(db, advertsCollection),
	}
}
