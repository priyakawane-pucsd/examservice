package questionservice

import (
	"context"
	"examservice/models/dao"
	"examservice/models/dto"
	"examservice/models/filters"
	"examservice/utils"
	"fmt"
	"net/http"
)

type Service struct {
	conf *Config
	repo Repository
}

type Config struct{}

type Repository interface {
	CreateOrUpdateQuestions(ctx context.Context, cfg *dao.Question) error
	GetQuestionsList(ctx context.Context, filter *filters.QuestionFilter, limit, offset int) ([]*dao.Question, error)
	GetQuestionById(ctx context.Context, questionId string) (*dao.Question, error)
	DeleteQuestionById(ctx context.Context, id string) error
	GetQuestionByUserId(ctx context.Context, id string, userId int64) error
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateOrUpdateQuestions(ctx context.Context, req *dto.QuestionRequest, questionId string) (string, error) {
	if questionId != "{id}" && questionId != "undefined" {
		_, err := s.repo.GetQuestionById(ctx, questionId)
		if err != nil {
			return "", err
		}
		req.ID = questionId
	}

	question := req.ToMongoObject()
	err := s.repo.CreateOrUpdateQuestions(ctx, question)
	if err != nil {
		return "", err
	}
	return "CreateOrUpdate Successfully", nil
}

func (s *Service) GetQuestionsList(ctx context.Context, filter *filters.QuestionFilter, limit, offset int) (*dto.ListQuestionResponse, error) {
	questions, err := s.repo.GetQuestionsList(ctx, filter, limit, offset)
	if err != nil {
		return nil, err
	}
	fmt.Println("questionquestion", questions)
	response := &dto.ListQuestionResponse{
		Questions: ConvertToQuestionResponseList(questions),
	}

	return response, nil
}

func (s *Service) GetQuestionById(ctx context.Context, questionId string, userId int64) (*dto.QuestionByIdResponse, error) {
	question, err := s.repo.GetQuestionById(ctx, questionId)
	if err != nil {
		return nil, err
	}
	if question.CreatedBy != userId {
		return nil, utils.NewUnauthorizedError("permission denied to delete question")
	}
	fmt.Println("questionquestion", question)

	response := &dto.QuestionByIdResponse{
		Question:   *QuestionsResponse(question),
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) DeleteQuestionById(ctx context.Context, id string, userId int64) error {
	err := s.repo.GetQuestionByUserId(ctx, id, userId)
	if err != nil {
		return err
	}

	err = s.repo.DeleteQuestionById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
