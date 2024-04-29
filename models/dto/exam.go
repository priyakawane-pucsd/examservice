package dto

import (
	"errors"
	"examservice/models/dao"
	"strings"
	"time"
)

type ExamRequest struct {
	ID              string   `json:"-"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	StartTime       int64    `json:"startTime"`
	EndTime         int64    `json:"endTime"`
	Duration        int      `json:"duration"`
	Questions       []string `json:"questions"`
	Topic           string   `json:"topic"`
	SubTopic        string   `json:"subTopic"`
	ExamFee         float64  `json:"examFee"`
	DifficultyLevel string   `json:"difficultyLevel"`
	CreatedBy       int64    `json:"-"`
}

var validDifficultLevels = map[string]bool{
	"EASY":   true,
	"MEDIUM": true,
	"HARD":   true,
}

func (er *ExamRequest) Validate() error {
	if !validDifficultLevels[strings.ToUpper(er.DifficultyLevel)] {
		return errors.New("invalid difficult level")
	}

	// Check if title is empty
	if er.Title == "" {
		return errors.New("title cannot be empty")
	}

	// Check if description is empty
	if er.Description == "" {
		return errors.New("description cannot be empty")
	}

	//check if any question is empty
	if len(er.Questions) == 0 {
		return errors.New("questions cannot be empty")
	}
	for _, question := range er.Questions {
		if question == "" {
			return errors.New("each question must not be empty")
		}
	}

	// Check if StartTime is a valid time format
	if er.StartTime <= 0 {
		return errors.New(" invalid start time")
	}

	// Check if EndTime is a valid time format
	if er.EndTime <= 0 {
		return errors.New(" invalid end time")
	}

	// Check if startTime must be less than endTime
	if er.EndTime <= er.StartTime {
		return errors.New(" startTime must be less than endTime")
	}

	// Check if Duration is non-negative
	if er.Duration <= 0 {
		return errors.New("duration cannot be negative")
	}

	// If all validations pass, return nil
	return nil
}

type ExamResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type Exam struct {
	ID              string   `json:"_id,omitempty"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	StartTime       int64    `json:"startTime"`
	EndTime         int64    `json:"endTime"`
	Duration        int      `json:"duration"`
	Questions       []string `json:"questions"`
	Topic           string   `json:"topic"`
	SubTopic        string   `json:"subTopic"`
	ExamFee         float64  `json:"examFee"`
	DifficultyLevel string   `json:"difficultyLevel"`
	IsDeleted       bool     `json:"isDeleted"`
	CreatedBy       int64    `json:"createdBy"`
	CreatedAt       int64    `json:"createdAt,omitempty"`
	UpdatedAt       int64    `json:"updatedAt,omitempty"`
}

func (r *ExamRequest) ToMongoObject() *dao.Exam {
	return &dao.Exam{
		ID:              r.ID,
		Title:           r.Title,
		Description:     r.Description,
		StartTime:       time.Now().UnixMilli(),
		EndTime:         time.Now().UnixMilli(),
		Duration:        r.Duration,
		Questions:       r.Questions,
		Topic:           r.Topic,
		SubTopic:        r.SubTopic,
		ExamFee:         r.ExamFee,
		CreatedBy:       r.CreatedBy,
		DifficultyLevel: r.DifficultyLevel,
		CreatedAt:       time.Now().UnixMilli(),
		UpdatedAt:       time.Now().UnixMilli(),
	}
}

type ListExamsResponse struct {
	Exams []Exam `json:"exam"`
}

type DeleteExamResponse struct {
	Message    string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type ExamByIdResponse struct {
	Exam       Exam `json:"exam"`
	StatusCode int  `json:"statusCode"`
}
