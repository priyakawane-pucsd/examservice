package mongo

import (
	"context"
	"examservice/models/dao"
	"examservice/utils"
	"time"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	ANSWER_COLLECTION = "answers"
)

func (r *Repository) CreateOrUpdateAnswer(ctx context.Context, req *dao.Answer) (string, error) {
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
		return "", utils.NewInternalServerError("Failed to upsert question into the database")
	}

	return objectID, nil
}
