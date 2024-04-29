package dto

import (
	"errors"
	"examservice/models/dao"
	"time"
)

type AnswerRequest struct {
	ID      string           `json:"-"`
	UserID  int64            `json:"-"`
	ExamID  string           `json:"examId"`
	Answers []QuestionAnswer `json:"answers"`
}

type Result struct {
	Attempted int64 `json:"attempted"`
	Correct   int64 `json:"correct"`
}

type QuestionAnswer struct {
	QuestionId string `json:"questionId"`
	Answer     string `json:"answer"`
}

type Answer struct {
	ID        string           `json:"_id,omitempty"`
	UserID    int64            `json:"userId"`
	ExamID    string           `json:"examId"`
	Answers   []QuestionAnswer `json:"answers"`
	Result    Result           `json:"result"`
	CreatedAt int64            `json:"createdAt"`
	UpdatedAt int64            `json:"updatedAt"`
}

type AnswerResponse struct {
	Id         string `json:"_id,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func (ar *AnswerRequest) ToMongoObject() *dao.Answer {
	var answers []dao.QuestionAnswer

	for _, ans := range ar.Answers {
		answers = append(answers, dao.QuestionAnswer{
			QuestionId: ans.QuestionId,
			Answer:     ans.Answer,
		})
	}

	return &dao.Answer{
		ID:        ar.ID,
		UserID:    ar.UserID,
		ExamID:    ar.ExamID,
		Answers:   answers,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
}

func (a *AnswerRequest) Validate() error {
	if a.ExamID == "" {
		return errors.New("examId is required")
	}

	if len(a.Answers) == 0 {
		return errors.New("answers must not be empty")
	}

	for _, ans := range a.Answers {
		if ans.Answer == "" {
			return errors.New("answer cannot be empty")
		}

		if ans.QuestionId == "" {
			return errors.New("answer's questionId cannot be empty")
		}
	}
	return nil
}
