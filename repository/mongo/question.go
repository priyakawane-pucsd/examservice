package mongo

import (
	"context"
	"errors"
	"examservice/models/dao"
	"examservice/models/filters"
	"examservice/utils"
	"time"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	QUESTIONS_COLLECTION = "questions"
)

func (r *Repository) CreateOrUpdateQuestions(ctx context.Context, req *dao.Question) (string, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Convert created_at and updated_at to milliseconds
	req.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	req.UpdatedAt = req.CreatedAt //Assume created_at and updated_at are the same initially

	// Set the ID of the question
	objectID := ""
	if req.ID == "" {
		objectID = primitive.NewObjectID().Hex()
	} else {
		objectID = req.ID
	}

	// Upsert the question document into the collection
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

func (r *Repository) GetQuestionsList(ctx context.Context, filter *filters.QuestionFilter, limit, offset int) ([]*dao.Question, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Define the filter based on provided topic and subTopic
	QueryFilter := bson.M{}
	if filter.Topic != "" {
		QueryFilter["topic"] = filter.Topic
	}
	if filter.SubTopic != "" {
		QueryFilter["subTopic"] = filter.SubTopic
	}
	if filter.UserId != "" {
		QueryFilter["userId"] = filter.UserId
	}

	// Define options for the find operation
	findOptions := options.Find().SetSort(bson.M{"createdAt": -1}).SetLimit(int64(limit)).SetSkip(int64(offset))

	// Execute the find operation to retrieve all questions
	cursor, err := collection.Find(ctx, QueryFilter, findOptions)
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

func (r *Repository) GetQuestionById(ctx context.Context, id string) (*dao.Question, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)
	filter := bson.M{"_id": id}
	var result *dao.Question
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NewCustomError(404, "Question not found")
		}
		logger.Error(ctx, "Error finding question by ID: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve question by ID")
	}
	return result, nil
}

func (r *Repository) GetQuestionsCountByIds(ctx context.Context, questionIds []string) (int64, error) {
	// Get the question collection
	questionCollection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	filter := bson.M{"_id": bson.M{"$in": questionIds}}
	count, err := questionCollection.CountDocuments(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error counting documents in questions collection: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *Repository) GetQuestionsByFilters(ctx context.Context, questionIds []string) ([]*dao.Question, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	// Define the filter based on provided topic and subTopic
	filter := bson.M{"_id": bson.M{"$in": questionIds}}

	// Define options for the find operation
	findOptions := options.Find()

	// Execute the find operation to retrieve all questions
	cursor, err := collection.Find(ctx, filter, findOptions)
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
