package questionservice

import (
	"context"
	"examservice/models/dao"
	"examservice/models/dto"
	"net/http"
)

type Service struct {
	conf *Config
	repo Repository
}

type Config struct{}

type Repository interface {
	CreateOrUpdateQuestions(ctx context.Context, cfg *dao.Question) (string, error)
	GetQuestionsList(ctx context.Context) ([]*dao.Question, error)
	DeleteQuestionById(ctx context.Context, id string) error
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateOrUpdateQuestions(ctx context.Context, req *dto.QuestionRequest) (*dto.QuestionResponse, error) {
	cfg := req.ToMongoObject()
	objectId, err := s.repo.CreateOrUpdateQuestions(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.QuestionResponse{StatusCode: http.StatusCreated, Id: objectId}, nil
}

func (s *Service) GetQuestionsList(ctx context.Context) (*dto.ListQuestionResponse, error) {
	questions, err := s.repo.GetQuestionsList(ctx)
	if err != nil {
		return nil, err
	}
	response := &dto.ListQuestionResponse{
		Questions:  ConvertToQuestionResponseList(questions),
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
