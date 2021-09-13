package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeedbackMongo struct {
	db *mongo.Collection
}

func (r *FeedbackMongo) AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error {
	objId, _ := primitive.ObjectIDFromHex(advertId)

	_, err := r.db.UpdateByID(ctx, objId, bson.D{{"$push", bson.D{{"feedbacks", feedback}}}})

	if err != nil {
		return err
	}

	return nil
}

func (r *FeedbackMongo) UpdateRating(ctx context.Context, advertId string) error {

	objId, _ := primitive.ObjectIDFromHex(advertId)

	var avgFeedback []bson.M

	matchStage := bson.D{{"$match", bson.D{{"_id", objId}}}}
	projectStage := bson.D{{"$project", bson.D{{"avgRating", bson.D{{"$avg", "$feedbacks.rating"}}}}}}

	cur, err := r.db.Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
	if err != nil {
		return err
	}

	if err = cur.All(ctx, &avgFeedback); err != nil {
		return err
	}

	averageRating := avgFeedback[0]["avgRating"]

	_, err = r.db.UpdateByID(ctx, objId, bson.M{ "$set": bson.M{"total_rating": averageRating}})
	if err != nil {
		return err
	}

	return nil
}

func NewFeedbackMongo(db *mongo.Database, collection string) *FeedbackMongo {
	return &FeedbackMongo{
		db: db.Collection(collection),
	}
}
