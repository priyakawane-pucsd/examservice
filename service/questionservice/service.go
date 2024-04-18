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
	CreateQuestions(ctx context.Context, cfg *dao.Question) error
	GetQuestionsList(ctx context.Context) ([]*dao.Question, error)
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateQuestions(ctx context.Context, req *dto.QuestionRequest) (*dto.QuestionResponse, error) {
	cfg := req.ToMongoObject()
	err := s.repo.CreateQuestions(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.QuestionResponse{StatusCode: http.StatusCreated, Id: cfg.ID}, nil
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
