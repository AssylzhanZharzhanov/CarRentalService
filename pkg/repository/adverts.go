package repository

import (
	"context"
	"fmt"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type AdvertMongo struct {
	db *mongo.Database
}
func (r *AdvertMongo) CreateAdvert(ctx context.Context, advert models.AdvertInput) (string, error) {
	res, err := r.db.Collection(advertsCollection).InsertOne(ctx, advert)

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AdvertMongo) GetMyAdverts(ctx context.Context, userId string) ([]models.Advert, error) {
	userObjId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"user_id": userObjId}

	adverts := make([]models.Advert, 0)

	cur, err := r.db.Collection(advertsCollection).Find(ctx, filter)
	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, nil
}

func (r *AdvertMongo) GetUserAdverts(ctx context.Context, userId string, status string) ([]models.Advert, error) {
	userObjId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"user_id": userObjId, "status": status}

	adverts := make([]models.Advert, 0)

	cur, err := r.db.Collection(advertsCollection).Find(ctx, filter)
	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, nil
}

func (r *AdvertMongo) GetTopAdverts(ctx context.Context) ([]models.Advert, error) {
	adverts := make([]models.Advert, 0)

	filter := bson.M{

	}

	cur, err := r.db.Collection(advertsCollection).Find(ctx, filter)
	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, nil}


func (r *AdvertMongo) GetSimilarAdverts(ctx context.Context, title string, price int) ([]models.Advert, error) {
	adverts := make([]models.Advert, 0)

	searchValue := fmt.Sprintf("\"%s\"", strings.ToLower(title))

	cur, err := r.db.Collection(advertsCollection).Find(ctx, bson.M{"$text": bson.M{"$search": searchValue, "$caseSensitive": false}, "price": bson.M{"$lte": price}})
	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, err
}

func (r *AdvertMongo) GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error) {
	adverts := make([]models.Advert, 0)

	cur, err := r.db.Collection(advertsCollection).Find(ctx, filter)
	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, nil
}

func (r *AdvertMongo) GetAdvertById(ctx context.Context, id string) (models.Advert, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.Collection(advertsCollection).UpdateOne(ctx, bson.D{{"_id", objId}},
		bson.D{{"$inc", bson.D{{"views", 1}}}})

	advert := models.Advert{}
	err = r.db.Collection(advertsCollection).FindOne(ctx, bson.M{"_id": objId}).Decode(&advert)
	if err != nil {
		return advert, err
	}

	return advert, nil
}

func (r *AdvertMongo) UpdateAdvert(ctx context.Context, id string, advert models.AdvertInput) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.Collection(advertsCollection).UpdateOne(ctx, bson.D{{"_id", objId}}, bson.M{"$set": advert})
	if err != nil {
		return err
	}

	return nil
}

func (r *AdvertMongo) DeleteAdvert(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.Collection(advertsCollection).DeleteOne(ctx, bson.D{{"_id", objId}})
	if err != nil {
		return err
	}

	return nil
}

func (r *AdvertMongo) UploadImage(ctx context.Context, advertId string, urls []string) error {
	advertObjId, _ := primitive.ObjectIDFromHex(advertId)

	filter := bson.M{"_id": advertObjId}

	imageList := make([]models.Image, 0)

	for _, url := range urls {
		imageList = append(imageList, models.Image{
			ID:  primitive.NewObjectID(),
			Url: url,
		})
	}

	_, err := r.db.Collection(advertsCollection).UpdateOne(ctx, filter, bson.M{"$push": bson.M{"images": bson.M{"$each": imageList}}})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func NewAdvertMongo(db *mongo.Database) *AdvertMongo {
	return &AdvertMongo{
		db: db,
	}
}
