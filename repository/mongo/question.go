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
	QUESTIONS_COLLECTION = "questions"
)

func (r *Repository) CreateOrUpdateQuestions(ctx context.Context, req *dao.Question) (string, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Set the ID of the question
	objectID := ""
	if req.ID == "" {
		objectID = primitive.NewObjectID().Hex()
	} else {
		objectID = req.ID
	}

	question := bson.M{
		"id":          objectID,
		"text":        req.Text,
		"choices":     req.Choices,
		"correct":     req.Correct,
		"explanation": req.Explanation,
		"userId":      req.UserId,
		"createdAt":   time.Now(),
		"updatedAt":   time.Now(),
	}

	// Upsert the question document into the collection
	filter := bson.M{"id": objectID}
	update := bson.M{"$set": question}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.Error(ctx, "Error upserting question: %v", err)
		return "", utils.NewInternalServerError("Failed to upsert question into the database")
	}

	return objectID, nil
}

func (r *Repository) GetQuestionsList(ctx context.Context) ([]*dao.Question, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Define options for the find operation
	findOptions := options.Find().SetSort(bson.M{"createdAt": -1})

	// Execute the find operation to retrieve all questions
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		logger.Error(ctx, "Error retrieving questions: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve questions from the database")
	}
	defer cursor.Close(ctx)

	// Decode all documents into a slice of dao.Question objects
	var questions []*dao.Question
	if err := cursor.All(ctx, &questions); err != nil {
		logger.Error(ctx, "Error decoding questions: %v", err)
		return nil, utils.NewInternalServerError("Failed to decode questions")
	}

	return questions, nil
}

func (r *Repository) DeleteQuestionById(ctx context.Context, id string) error {
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Specify the filter based on the ID
	filter := bson.M{"_id": id}
	// Perform the deletion operation
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error deleting question by ID: %v", err)
		return utils.NewInternalServerError("Failed to delete question")
	}

	// Check if no documents were matched and deleted
	if result.DeletedCount == 0 {
		return utils.NewCustomError(404, "Question not found with this Id")
	}

	return nil
}
