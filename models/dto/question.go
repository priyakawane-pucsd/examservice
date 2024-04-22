package dto

import (
	"errors"
	"examservice/models/dao"
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

type DeleteQuestionResponse struct {
	Message    string `json:"_id,omitempty"`
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
		// CreatedAt:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		// UpdatedAt:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// ListQuestionResponse represents the response format for the list of questions.
type ListQuestionResponse struct {
	Questions  []Question `json:"questions"`
	StatusCode int        `json:"statusCode"`
}

// Question represents a single question.
type Question struct {
	ID          string    `json:"_id"`
	Text        string    `json:"text"`
	Choices     []string  `json:"choices"`
	Correct     string    `json:"correct"`
	Explanation string    `json:"explanation"`
	UserId      string    `json:"userId"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	// CreatedAt   string   `json:"createdAt"`
	// UpdatedAt   string   `json:"updatedAt"`
}

func (q *QuestionRequest) Validate() error {
	if q.Text == "" {
		return errors.New("text is required")
	}
	if len(q.Choices) == 0 {
		return errors.New("choices must not be empty")
	}

	for _, choice := range q.Choices {
		if choice == "" {
			return errors.New("choice cannot be empty")
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
