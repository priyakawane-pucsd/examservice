package dto

import (
	"errors"
	"examservice/models/dao"
	"time"
)

type ExamRequest struct {
	ID          string   `json:"_id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
	Duration    int      `json:"duration"`
	Questions   []string `json:"questions"`
	Topic       string   `json:"topic"`
	SubTopic    string   `json:"sub_topic"`
}

func (er *ExamRequest) Validate() error {
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
	_, err := time.Parse(time.RFC3339, er.StartTime)
	if err != nil {
		return errors.New("start_time must be a valid RFC3339 formatted time")
	}
	// Check if EndTime is a valid time format
	_, err = time.Parse(time.RFC3339, er.EndTime)
	if err != nil {
		return errors.New("end_time must be a valid RFC3339 formatted time")
	}
	// Check if Duration is non-negative
	if er.Duration < 0 {
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
	ID          string    `json:"_id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	Duration    int       `json:"duration"`
	Questions   []string  `json:"questions"`
	Topic       string    `json:"topic"`
	SubTopic    string    `json:"sub_topic"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (r *ExamRequest) ToMongoObject() *dao.Exam {
	return &dao.Exam{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		StartTime:   r.StartTime,
		EndTime:     r.EndTime,
		Duration:    r.Duration,
		Questions:   r.Questions,
		Topic:       r.Topic,
		SubTopic:    r.SubTopic,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		// CreatedAt:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		// UpdatedAt:   strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
}

type ListExamsResponse struct {
	Exams      []Exam `json:"exam"`
	StatusCode int    `json:"statusCode"`
}

type DeleteExamResponse struct {
	Message    string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}
