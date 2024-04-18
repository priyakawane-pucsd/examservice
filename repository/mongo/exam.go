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
	EXAM_COLLECTION = "exams"
)

func (r *Repository) CreateExam(ctx context.Context, req *dao.Exam) error {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)

	// Set the ID of the exam
	objectId := primitive.NewObjectID()

	exam := bson.M{
		"_id":         objectId,
		"title":       req.Title,
		"description": req.Description,
		"startTime":   req.StartTime,
		"duration":    req.Duration,
		"questions":   req.Questions,
		"createdAt":   time.Now().Format(time.RFC3339),
		"updatedAt":   time.Now().Format(time.RFC3339),
	}
	// Insert the exam document into the collection
	_, err := collection.InsertOne(ctx, exam)
	if err != nil {
		logger.Error(ctx, "Error inserting new exam: %v", err)
		return utils.NewInternalServerError("Failed to insert new exam into the database")
	}
	return nil
}

func (r *Repository) GetExamsList(ctx context.Context) ([]*dao.Exam, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	// Define options for the find operation
	findOptions := options.Find()

	// Execute the find operation to retrieve all questions
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		logger.Error(ctx, "Error retrieving exams: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve exams from the database")
	}
	defer cursor.Close(ctx)

	// Decode all documents into a slice of dao.Question objects
	var exams []*dao.Exam
	if err := cursor.All(ctx, &exams); err != nil {
		logger.Error(ctx, "Error decoding exams: %v", err)
		return nil, utils.NewInternalServerError("Failed to decode exams")
	}

	return exams, nil
}
