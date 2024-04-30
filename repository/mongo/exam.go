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
	EXAM_COLLECTION = "exams"
)

func (r *Repository) CreateOrUpdateExam(ctx context.Context, req *dao.Exam) error {
	// Specify the MongoDB collection
	examCollection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)

	// Convert created_at and updated_at to milliseconds
	createdAtMillis := time.Now().UnixNano() / int64(time.Millisecond)
	updatedAtMillis := createdAtMillis // Assume created_at and updated_at are the same initially

	req.CreatedAt = createdAtMillis
	req.UpdatedAt = updatedAtMillis

	// Set the ID of the exam
	objectId := ""
	if req.ID == "" {
		objectId = primitive.NewObjectID().Hex()
	} else {
		objectId = req.ID
	}

	//Upsert the exam document into the collection
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": req}

	_, err := examCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		logger.Error(ctx, "Error upserting exam: %v", err)
		return utils.NewInternalServerError("Failed to upsert exam into the database")
	}
	return nil
}

func (r *Repository) GetExamsList(ctx context.Context, filter *filters.ExamFilter, limit, offset int) ([]*dao.Exam, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)

	// Define the filter based on provided topic and subTopic
	QueryFilter := bson.M{
		"isDeleted": bson.M{"$ne": true}, // Exclude deleted questions
	}
	if filter.Topic != "" {
		QueryFilter["topic"] = filter.Topic
	}
	if filter.SubTopic != "" {
		QueryFilter["sub_topic"] = filter.SubTopic
	}
	QueryFilter["createdBy"] = filter.UserId

	// Define options for the find operation
	findOptions := options.Find().SetSort(bson.M{"createdAt": -1}).SetLimit(int64(limit)).SetSkip(int64(offset))

	// Execute the find operation to retrieve all questions
	cursor, err := collection.Find(ctx, QueryFilter, findOptions)
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

func (r *Repository) GetExamById(ctx context.Context, id string) (*dao.Exam, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	filter := bson.M{"_id": id}
	var result *dao.Exam
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NewCustomError(404, "Exam not found")
		}
		logger.Error(ctx, "Error finding exam by ID: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve exam by ID")
	}
	return result, nil
}

func (r *Repository) DeleteExamById(ctx context.Context, id string) error {
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isDeleted": true}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error(ctx, "Error soft deleting exam by ID %v", err)
		return utils.NewInternalServerError("Failed to soft delete exam")
	}
	if res.ModifiedCount == 0 {
		utils.NewCustomError(404, "No exam found for this id")
	}
	return nil
}

func (r *Repository) GetExamByUserId(ctx context.Context, id string, userId int64) error {
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	filter := bson.M{"_id": id}
	var result *dao.Exam
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return utils.NewCustomError(404, "Exam not found")
		}
		logger.Error(ctx, "Error finding exam by ID: %v", err)
		return utils.NewInternalServerError("Failed to retrieve exam by ID")
	}
	if result.CreatedBy != userId {
		return utils.NewUnauthorizedError("Permission denied to access this question")
	}
	return nil
}
