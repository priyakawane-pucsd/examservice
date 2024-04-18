package examservice

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
	CreateExam(ctx context.Context, cfg *dao.Exam) error
	GetExamsList(ctx context.Context) ([]*dao.Exam, error)
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateExam(ctx context.Context, req *dto.ExamRequest) (*dto.ExamResponse, error) {
	cfg := req.ToMongoObject()
	err := s.repo.CreateExam(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.ExamResponse{StatusCode: http.StatusCreated, Id: cfg.ID}, nil
}

func (s *Service) GetExamsList(ctx context.Context) (*dto.ListExamsResponse, error) {
	exams, err := s.repo.GetExamsList(ctx)
	if err != nil {
		return nil, err
	}
	response := &dto.ListExamsResponse{
		Exams:      ConvertToExamResponseList(exams),
		StatusCode: http.StatusOK,
	}
	return response, nil
}
