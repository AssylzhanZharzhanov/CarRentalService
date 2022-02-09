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
	feedback.Id = primitive.NewObjectID()
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

func (r *FeedbackMongo) UpdateFeedback(ctx context.Context, feedbackId string, feedback models.Feedback) error {
	panic("implement me")
}

func (r *FeedbackMongo) DeleteFeedback(ctx context.Context, feedbackId string) error {
	panic("implement me")
}

func NewFeedbackMongo(db *mongo.Database) *FeedbackMongo {
	return &FeedbackMongo{
		db: db.Collection(advertsCollection),
	}
}
