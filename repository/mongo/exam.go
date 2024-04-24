package mongo

import (
	"context"
	"errors"
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

func (r *Repository) validateQuestionIds(ctx context.Context, questionIds []string) (bool, error) {
	// Get the question collection
	questionCollection := r.conn.Database(r.cfg.Database).Collection(QUESTIONS_COLLECTION)

	filter := bson.M{"_id": bson.M{"$in": questionIds}}
	count, err := questionCollection.CountDocuments(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error counting documents in questions collection: %v", err)
		return false, err
	}
	if count != int64(len(questionIds)) {
		return false, nil // At least one ID doesn't exist
	}

	// Iterate over the question IDs and check if each ID exists in the database
	// for _, id := range questionIds {
	// 	filter := bson.M{"_id": id}
	// 	count, err := questionCollection.CountDocuments(ctx, filter)
	// 	if err != nil {
	// 		logger.Error(ctx, "Error counting documents in questions collection: %v", err)
	// 		return false, err
	// 	}
	// 	if count == 0 {
	// 		return false, nil // At least one ID doesn't exist
	// 	}
	// }
	return true, nil
}

func (r *Repository) CreateOrUpdateExam(ctx context.Context, req *dao.Exam) (string, error) {
	// Specify the MongoDB collection
	examCollection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)

	// Validate question IDs if provided
	if len(req.Questions) != 0 {
		valid, err := r.validateQuestionIds(ctx, req.Questions)
		if err != nil {
			logger.Error(ctx, "Error validating question IDs: %v", err)
			return "", err
		}
		if !valid {
			return "", errors.New("one or more question IDs are invalid")
		}
	}

	// Convert created_at and updated_at to milliseconds
	createdAtMillis := time.Now().UnixNano() / int64(time.Millisecond)
	updatedAtMillis := createdAtMillis // Assume created_at and updated_at are the same initially

	exam := bson.M{
		"title":           req.Title,
		"description":     req.Description,
		"startTime":       req.StartTime,
		"endTime":         req.EndTime,
		"duration":        req.Duration,
		"topic":           req.Topic,
		"sub_topic":       req.SubTopic,
		"questions":       req.Questions,
		"examFee":         req.ExamFee,
		"difficultyLevel": req.DifficultyLevel,
		"createdAt":       createdAtMillis,
		"updatedAt":       updatedAtMillis,
	}

	// Set the ID of the exam
	objectId := ""
	if req.ID == "" {
		objectId = primitive.NewObjectID().Hex()
	} else {
		objectId = req.ID
	}

	//Upsert the exam document into the collection
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": exam}

	_, err := examCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		logger.Error(ctx, "Error upserting exam: %v", err)
		return "", utils.NewInternalServerError("Failed to upsert exam into the database")
	}
	return objectId, nil
}

func (r *Repository) GetExamsList(ctx context.Context, topic string, subTopic string) ([]*dao.Exam, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)

	// Define the filter based on provided topic and subTopic
	filter := bson.M{}
	if topic != "" {
		filter["topic"] = bson.M{"$regex": primitive.Regex{Pattern: topic, Options: "i"}}
	}
	if subTopic != "" {
		filter["sub_topic"] = bson.M{"$regex": primitive.Regex{Pattern: subTopic, Options: "i"}}
	}

	// Define options for the find operation
	findOptions := options.Find().SetSort(bson.M{"createdAt": -1})

	// Execute the find operation to retrieve all questions
	cursor, err := collection.Find(ctx, filter, findOptions)
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

func (r *Repository) DeleteExamById(ctx context.Context, id string) error {
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error deleting exam by ID: %v", err)
		return utils.NewInternalServerError("Failed to delete exam")
	}
	if res.DeletedCount == 0 {
		utils.NewCustomError(404, "No exam found for this id")
	}
	return nil
}
