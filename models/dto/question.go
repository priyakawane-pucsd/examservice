package dto

import (
	"errors"
	"examservice/models/dao"
	"time"
)

type QuestionRequest struct {
	ID          string   `json:"-"`
	Text        string   `json:"text"`
	Choices     []Choice `json:"choices"`
	Correct     string   `json:"correct"`
	Explanation string   `json:"explanation"`
	Topic       string   `json:"topic"`
	SubTopic    string   `json:"subTopic"`
	CreatedBy   int64    `json:"-"`
}

type CreateQuestionResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type QuestionResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type QuestionError struct {
	Error      string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type DeleteQuestionResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (r *QuestionRequest) ToMongoObject() *dao.Question {
	var choices []dao.Choice
	for _, choice := range r.Choices {
		choices = append(choices, dao.Choice{
			Key:   choice.Key,
			Value: choice.Value,
		})
	}

	return &dao.Question{
		ID:          r.ID,
		Text:        r.Text,
		Choices:     choices,
		Correct:     r.Correct,
		Explanation: r.Explanation,
		Topic:       r.Topic,
		SubTopic:    r.SubTopic,
		CreatedBy:   r.CreatedBy,
		CreatedAt:   time.Now().UnixMilli(),
		UpdatedAt:   time.Now().UnixMilli(),
	}
}

// ListQuestionResponse represents the response format for the list of questions.
type ListQuestionResponse struct {
	Questions []Question `json:"questions"`
}

type Choice struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Question represents a single question.
type Question struct {
	ID          string   `json:"_id"`
	Text        string   `json:"text"`
	Choices     []Choice `json:"choices"`
	Correct     string   `json:"correct"`
	Explanation string   `json:"explanation"`
	Topic       string   `json:"topic"`
	SubTopic    string   `json:"subTopic"`
	CreatedBy   int64    `json:"createdBy"`
	IsDeleted   bool     `json:"isDeleted"`
	CreatedAt   int64    `json:"createdAt,omitempty"`
	UpdatedAt   int64    `json:"updatedAt,omitempty"`
}

func (q *QuestionRequest) Validate() error {
	if q.Text == "" {
		return errors.New("text is required")
	}
	if len(q.Choices) == 0 {
		return errors.New("choices must not be empty")
	}

	for _, choice := range q.Choices {
		if choice.Key == "" {
			return errors.New("choice key cannot be empty")
		}

		if choice.Value == "" {
			return errors.New("choice value cannot be empty")
		}
	}
	if q.Correct == "" {
		return errors.New("correct answer is required")
	}

	if q.Explanation == "" {
		return errors.New("explanation is required")
	}
	return nil
}

type QuestionByIdResponse struct {
	Question   Question `json:"question"`
	StatusCode int      `json:"statusCode"`
}
