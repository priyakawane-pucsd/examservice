package mongo

import (
	"context"
	"errors"
	"examservice/models/dao"
	"examservice/utils"
	"time"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	ANSWER_COLLECTION = "answers"
)

func (r *Repository) CreateOrUpdateAnswer(ctx context.Context, req *dao.Answer) error {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(ANSWER_COLLECTION)

	createdAtMillis := time.Now().UnixNano() / int64(time.Millisecond)
	updatedAtMillis := createdAtMillis

	req.CreatedAt = createdAtMillis
	req.UpdatedAt = updatedAtMillis

	objectID := ""
	if req.ID == "" {
		objectID = primitive.NewObjectID().Hex()
	} else {
		objectID = req.ID
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": req}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.Error(ctx, "Error upserting question: %v", err)
		return utils.NewInternalServerError("Failed to upsert question into the database")
	}

	return nil
}

func (r *Repository) GetAnswerById(ctx context.Context, id string) (*dao.Exam, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(ANSWER_COLLECTION)
	filter := bson.M{"_id": id}
	var result *dao.Exam
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NewCustomError(404, "Answer not found")
		}
		logger.Error(ctx, "Error finding answer by ID: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve answer by ID")
	}
	return result, nil
}
