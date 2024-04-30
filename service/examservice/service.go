package examservice

import (
	"context"
	"examservice/models/dao"
	"examservice/models/dto"
	"examservice/models/filters"
	"examservice/utils"
	"net/http"
)

type Service struct {
	conf *Config
	repo Repository
}

type Config struct{}

type Repository interface {
	CreateOrUpdateExam(ctx context.Context, cfg *dao.Exam) error
	GetExamsList(ctx context.Context, filter *filters.ExamFilter, limit, offset int) ([]*dao.Exam, error)
	GetExamById(ctx context.Context, examId string) (*dao.Exam, error)
	DeleteExamById(ctx context.Context, id string) error
	GetQuestionsCountByIds(ctx context.Context, questionIds []string) (int64, error)
	GetExamByUserId(ctx context.Context, id string, userId int64) error
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) CreateOrUpdateExam(ctx context.Context, req *dto.ExamRequest, examId string) (string, error) {
	if examId != "{id}" && examId != "undefined" {
		_, err := s.repo.GetExamById(ctx, examId)
		if err != nil {
			return "", err
		}
		req.ID = examId
	}

	dbQuestionCount, err := s.repo.GetQuestionsCountByIds(ctx, req.Questions)
	if err != nil {
		return "", err
	}
	if len(req.Questions) != int(dbQuestionCount) {
		return "", utils.NewBadRequestError("Invalid Question Ids")
	}

	exam := req.ToMongoObject()
	err = s.repo.CreateOrUpdateExam(ctx, exam)
	if err != nil {
		return "", err
	}
	return "CreateOrUpdate Successfully", nil
}

func (s *Service) GetExamsList(ctx context.Context, filter *filters.ExamFilter, limit, offset int) (*dto.ListExamsResponse, error) {
	exams, err := s.repo.GetExamsList(ctx, filter, limit, offset)
	if err != nil {
		return nil, err
	}
	response := &dto.ListExamsResponse{
		Exams: ConvertToExamResponseList(exams),
	}
	return response, nil
}

func (s *Service) GetExamById(ctx context.Context, examId string, userId int64) (*dto.ExamByIdResponse, error) {
	exam, err := s.repo.GetExamById(ctx, examId)
	if err != nil {
		return nil, err
	}
	if exam.CreatedBy != userId {
		return nil, utils.NewUnauthorizedError("Permission denied to access this question")
	}
	response := &dto.ExamByIdResponse{
		Exam:       *ExamsResponse(exam),
		StatusCode: http.StatusOK,
	}
	return response, nil
}

func (s *Service) DeleteExamById(ctx context.Context, examId string, userId int64) error {
	err := s.repo.GetExamByUserId(ctx, examId, userId)
	if err != nil {
		return err
	}

	err = s.repo.DeleteExamById(ctx, examId)
	if err != nil {
		return err
	}
	return nil
}
