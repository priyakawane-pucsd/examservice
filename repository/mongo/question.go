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

func (r *Repository) CreateQuestions(ctx context.Context, req *dao.Question) error {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Generate a new ObjectID for the question
	objectID := primitive.NewObjectID()

	question := bson.M{
		"_id":         objectID,
		"text":        req.Text,
		"choices":     req.Choices,
		"correct":     req.Correct,
		"explanation": req.Explanation,
		"userId":      req.UserId,
		"createdAt":   time.Now().Format(time.RFC3339),
		"updatedAt":   time.Now().Format(time.RFC3339),
	}

	// Insert the document into the collection
	_, err := collection.InsertOne(ctx, question)
	if err != nil {
		logger.Error(ctx, "Error inserting new question: %v", err)
		return utils.NewInternalServerError("Failed to insert new question into the database")
	}
	return nil
}

func (r *Repository) GetQuestionsList(ctx context.Context) ([]*dao.Question, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Define options for the find operation
	findOptions := options.Find()

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
