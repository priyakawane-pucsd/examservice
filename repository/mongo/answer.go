package mongo

import (
	"context"
	"errors"
	"examservice/models/dao"
	"examservice/utils"
	"fmt"
	"time"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	ANSWER_COLLECTION = "answers"
)

func (r *Repository) ValidateExamId(ctx context.Context, examId string) (bool, error) {
	// Get the exam collection
	examCollection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	filter := bson.M{"_id": examId}

	count, err := examCollection.CountDocuments(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error counting documents in exam collection: %v", err)
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *Repository) CheckTimeOfSubmission(ctx context.Context, examId string) bool {
	// Get the exam collection
	examCollection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	filter := bson.M{"_id": examId}

	// Find the document with the given examId
	var result dao.Exam
	err := examCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		logger.Error(ctx, "Error finding document in exam collection: %v", err)
		return false
	}
	fmt.Println("RESULTTTTTT", result)

	// Get the current time
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	// Calculate the end time plus 5 minutes
	endTimePlus5 := result.EndTime + (5 * 60 * 1000) //5 minutes in milliseconds

	if currentTime >= result.EndTime && currentTime <= endTimePlus5 {
		// Submission is allowed

		return true
	}
	// Submission is not allowed
	return false
}

func (r *Repository) CreateOrUpdateAnswer(ctx context.Context, req *dao.Answer) (string, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(ANSWER_COLLECTION)

	//validate examId before insertion
	if req.ExamID != "" {
		valid, err := r.ValidateExamId(ctx, req.ExamID)
		if err != nil {
			logger.Error(ctx, "Error validating examId: %v", err)
			return "", err
		}
		if !valid {
			return "", errors.New("examId is invalid")
		}
	}

	createdAtMillis := time.Now().UnixNano() / int64(time.Millisecond)
	updatedAtMillis := createdAtMillis

	answer := bson.M{
		"examId":    req.ExamID,
		"userId":    req.UserID,
		"answers":   req.Answers,
		"createdAt": createdAtMillis,
		"updatedAt": updatedAtMillis,
	}

	objectID := ""
	if req.ID == "" {
		objectID = primitive.NewObjectID().Hex()
	} else {
		objectID = req.ID
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": answer}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.Error(ctx, "Error upserting question: %v", err)
		return "", utils.NewInternalServerError("Failed to upsert question into the database")
	}

	return objectID, nil
}
