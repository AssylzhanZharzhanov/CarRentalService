package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookmarkMongo struct {
	db *mongo.Database
}

func (r *BookmarkMongo) AddUserBookmark(ctx context.Context, userId string, advertId string) error {
	userObjId, _ := primitive.ObjectIDFromHex(userId)
	advertObjId, _ := primitive.ObjectIDFromHex(advertId)

	filter := bson.M{"_id": userObjId}
	_, err := r.db.Collection(usersCollection).UpdateOne(ctx, filter, bson.M{"addToSet": bson.M{"bookmarks": advertObjId}})

	return err
}

func (r *BookmarkMongo) GetUserBookmarks(ctx context.Context, userId string) ([]models.Advert, error) {
	userObjId, _ := primitive.ObjectIDFromHex(userId)

	var user models.User
	adverts := make([]models.Advert, 0)
	err := r.db.Collection(usersCollection).FindOne(ctx, bson.M{"_id": userObjId}).Decode(&user)

	cur, err := r.db.Collection(advertsCollection).Find(ctx, bson.M{"_id": bson.M{"$in": user.Bookmarks} })

	if err != nil {
		return adverts, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return adverts, err
	}

	return adverts, err
}

func NewBookmarkMongo(db *mongo.Database) *BookmarkMongo {
	return &BookmarkMongo{
		db: db,
	}
}
