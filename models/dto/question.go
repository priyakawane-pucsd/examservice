package dto

import (
	"errors"
	"examservice/models/dao"
	"strconv"
	"time"
)

type QuestionRequest struct {
	ID          string   `json:"_id,omitempty"`
	Text        string   `json:"text"`
	Choices     []string `json:"choices"`
	Correct     string   `json:"correct"`
	Explanation string   `json:"explanation"`
	UserId      string   `json:"userId"`
}

type QuestionResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func (r *QuestionRequest) ToMongoObject() *dao.Question {
	return &dao.Question{
		ID:          r.ID,
		Text:        r.Text,
		Choices:     r.Choices,
		Correct:     r.Correct,
		Explanation: r.Explanation,
		UserId:      r.UserId,
		CreatedAt:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		UpdatedAt:   strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
}

// ListQuestionResponse represents the response format for the list of questions.
type ListQuestionResponse struct {
	Questions  []Question `json:"questions"`
	StatusCode int        `json:"statusCode"`
}

// Question represents a single question.
type Question struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	Choices     []string `json:"choices"`
	Correct     string   `json:"correct"`
	Explanation string   `json:"explanation"`
	UserId      string   `json:"userId"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

func (q *QuestionRequest) Validate() error {
	if q.ID == "" {
		return errors.New("ID is required")
	}

	if q.Text == "" {
		return errors.New("Text is required")
	}

	if len(q.Choices) == 0 {
		return errors.New("Choices must not be empty")
	}
	for _, choice := range q.Choices {
		if choice == "" {
			return errors.New("Choice cannot be empty")
		}
	}

	if q.Correct == "" {
		return errors.New("Correct answer is required")
	}

	if q.Explanation == "" {
		return errors.New("Explanation is required")
	}

	if q.UserId == "" {
		return errors.New("UserID is required")
	}

	return nil
}
