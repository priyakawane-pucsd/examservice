package questionservice

import (
	"context"
	"examservice/models/dao"
	"examservice/models/dto"
	"examservice/models/filters"
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
	response := &dto.ListQuestionResponse{
		Questions:  ConvertToQuestionResponseList(questions),
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) GetQuestionById(ctx context.Context, questionId string) (*dto.QuestionByIdResponse, error) {
	question, err := s.repo.GetQuestionById(ctx, questionId)
	if err != nil {
		return nil, err
	}

	response := &dto.QuestionByIdResponse{
		Question:   *QuestionsResponse(question),
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) DeleteQuestionById(ctx context.Context, id string) (*dto.DeleteQuestionResponse, error) {
	err := s.repo.DeleteQuestionById(ctx, id)
	if err != nil {
		return nil, err
	}

	// If the deletion is successful, create a response
	response := &dto.DeleteQuestionResponse{
		Message:    "Question deleted successfully",
		StatusCode: http.StatusOK,
	}
	return response, nil
}
