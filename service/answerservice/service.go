package answerservice

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
	CreateOrUpdateAnswer(ctx context.Context, cfg *dao.Answer) (string, error)
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateOrUpdateAnswer(ctx context.Context, req *dto.AnswerRequest) (*dto.AnswerResponse, error) {
	cfg := req.ToMongoObject()
	objectId, err := s.repo.CreateOrUpdateAnswer(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.AnswerResponse{StatusCode: http.StatusCreated, Id: objectId}, nil
}
