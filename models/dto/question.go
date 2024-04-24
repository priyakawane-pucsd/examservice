package dto

import (
	"errors"
	"examservice/models/dao"
	"time"
)

type QuestionRequest struct {
	ID          string   `json:"_id,omitempty"`
	Text        string   `json:"text"`
	Choices     []Choice `json:"choices"`
	Correct     string   `json:"correct"`
	Explanation string   `json:"explanation"`
	UserId      string   `json:"userId"`
}

type QuestionResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

type DeleteQuestionResponse struct {
	Message    string `json:"_id,omitempty"`
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
		UserId:      r.UserId,
		CreatedAt:   time.Now().UnixMilli(),
		UpdatedAt:   time.Now().UnixMilli(),
	}
}

// ListQuestionResponse represents the response format for the list of questions.
type ListQuestionResponse struct {
	Questions  []Question `json:"questions"`
	StatusCode int        `json:"statusCode"`
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
	UserId      string   `json:"userId"`
	CreatedAt   int64    `json:"created_at,omitempty"`
	UpdatedAt   int64    `json:"updated_at,omitempty"`
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
	if q.UserId == "" {
		return errors.New("userID is required")
	}
	return nil
}

type QuestionByIdResponse struct {
	Question   Question `json:"question"`
	StatusCode int      `json:"statusCode"`
}
