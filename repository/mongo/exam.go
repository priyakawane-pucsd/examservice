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

func (r *Repository) CreateOrUpdateExam(ctx context.Context, req *dao.Exam) (string, error) {
	// Specify the MongoDB collection
	collection := r.conn.Database(r.cfg.Database).Collection(EXAM_COLLECTION)

	// Set the ID of the exam
	objectId := ""
	if req.ID == "" {
		objectId = primitive.NewObjectID().Hex()
	} else {
		objectId = req.ID
	}

	exam := bson.M{
		"title":       req.Title,
		"description": req.Description,
		"startTime":   req.StartTime,
		"endTime":     req.EndTime,
		"duration":    req.Duration,
		"topic":       req.Topic,
		"sub_topic":   req.SubTopic,
		"questions":   req.Questions,
		"createdAt":   time.Now().Format(time.RFC3339),
		"updatedAt":   time.Now().Format(time.RFC3339),
	}

	//Upsert the exam document into the collection
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": exam}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
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
