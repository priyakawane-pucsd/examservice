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
	CreateOrUpdateExam(ctx context.Context, cfg *dao.Exam) (string, error)
	GetExamsList(ctx context.Context, topic string, subTopic string) ([]*dao.Exam, error)
	DeleteExamById(ctx context.Context, id string) error
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateOrUpdateExam(ctx context.Context, req *dto.ExamRequest) (*dto.ExamResponse, error) {
	cfg := req.ToMongoObject()
	objectId, err := s.repo.CreateOrUpdateExam(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &dto.ExamResponse{StatusCode: http.StatusCreated, Id: objectId}, nil
}

func (s *Service) GetExamsList(ctx context.Context, topic string, subTopic string) (*dto.ListExamsResponse, error) {
	exams, err := s.repo.GetExamsList(ctx, topic, subTopic)
	if err != nil {
		return nil, err
	}
	response := &dto.ListExamsResponse{
		Exams:      ConvertToExamResponseList(exams),
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) DeleteExamById(ctx context.Context, examId string) (*dto.DeleteExamResponse, error) {
	err := s.repo.DeleteExamById(ctx, examId)
	if err != nil {
		return nil, err
	}
	// If the deletion is successful, create a response
	response := &dto.DeleteExamResponse{
		Message:    "Exam deleted successfully",
		StatusCode: http.StatusOK,
	}
	return response, nil
}
