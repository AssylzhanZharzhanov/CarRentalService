package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	categoriesCollection = "categories"
	citiesCollection = "cities"
	priceCollection = "prices"
	rentTypeCollection = "rent_types"
	statusCollection = "statuses"
)

type FilterRepository struct {
	db *mongo.Database
}

func (r *FilterRepository) AddStatus(ctx context.Context, status models.Status) error {
	_, err := r.db.Collection(statusCollection).InsertOne(ctx, status)
	return err
}

func (r *FilterRepository) GetStatuses(ctx context.Context) ([]models.Status, error) {
	statuses := make([]models.Status, 0)
	cur, err := r.db.Collection(statusCollection).Find(ctx, bson.M{})
	if err != nil {
		return statuses, err
	}

	if err = cur.All(ctx, &statuses); err != nil {
		return statuses, err
	}

	return statuses, nil
}

func (r *FilterRepository) DeleteStatus(ctx context.Context, name string) error {
	_, err := r.db.Collection(statusCollection).DeleteOne(ctx, bson.M{"name": name})
	return err
}


func (r *FilterRepository) AddRentType(ctx context.Context, rentType models.RentTypes) error {
	_, err := r.db.Collection(rentTypeCollection).InsertOne(ctx, rentType)
	return err
}

func (r *FilterRepository) GetRentTypes(ctx context.Context) ([]models.RentTypes, error) {
	rentTypes := make([]models.RentTypes, 0)
	cur, err := r.db.Collection(rentTypeCollection).Find(ctx, bson.M{})
	if err != nil {
		return rentTypes, err
	}

	if err = cur.All(ctx, &rentTypes); err != nil {
		return rentTypes, err
	}

	return rentTypes, nil
}

func (r *FilterRepository) DeleteRentType(ctx context.Context, name string) error {
	_, err := r.db.Collection(rentTypeCollection).DeleteOne(ctx, bson.M{"name": name})
	return err
}


func (r *FilterRepository) AddPrice(ctx context.Context, price models.Price) error {
	_, err := r.db.Collection(priceCollection).InsertOne(ctx, price)
	return err
}

func (r *FilterRepository) GetPrices(ctx context.Context) ([]models.Price, error) {
	prices := make([]models.Price, 0)
	cur, err := r.db.Collection(priceCollection).Find(ctx, bson.M{})
	if err != nil {
		return prices, err
	}

	if err = cur.All(ctx, &prices); err != nil {
		return prices, err
	}

	return prices, nil
}

func (r *FilterRepository) DeletePrices(ctx context.Context, name string) error {
	_, err := r.db.Collection(priceCollection).DeleteOne(ctx, bson.M{"name": name})
	return err
}


func (r *FilterRepository) AddCity(ctx context.Context, city models.City) error {
	_, err := r.db.Collection(citiesCollection).InsertOne(ctx, city)
	if err != nil {
		return err
	}
	return nil
}

func (r *FilterRepository) GetCities(ctx context.Context) ([]models.City, error) {
	cities := make([]models.City, 0)

	cur, err := r.db.Collection(citiesCollection).Find(ctx, bson.M{})
	if err != nil {
		return cities, err
	}

	if err = cur.All(ctx, &cities); err != nil {
		return cities, err
	}

	return cities, err
}

func (r *FilterRepository) DeleteCity(ctx context.Context, name string) error {
	_, err := r.db.Collection(citiesCollection).DeleteOne(ctx, bson.M{"name": name})
	return err
}


func (r *FilterRepository) AddCategory(ctx context.Context, category models.Category) error {
	_, err := r.db.Collection(categoriesCollection).InsertOne(ctx, category)
	return err
}

func (r *FilterRepository) GetCategories(ctx context.Context) ([]models.Category, error) {
	categories := make([]models.Category, 0)

	filter := bson.M{}
	cur, err := r.db.Collection(categoriesCollection).Find(ctx, filter)

	if err = cur.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *FilterRepository) DeleteCategory(ctx context.Context, name string) error {
	_, err := r.db.Collection(categoriesCollection).DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		return err
	}

	return err
}

func NewFilterRepository(db *mongo.Database) *FilterRepository {
	return &FilterRepository{db:db}
}
