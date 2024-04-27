package answerservice

import (
	"context"
	"errors"
	"examservice/models/dao"
	"examservice/models/dto"
	"net/http"
	"time"

	"github.com/bappaapp/goutils/logger"
)

type Service struct {
	conf *Config
	repo Repository
}

type Config struct{}

type Repository interface {
	CreateOrUpdateAnswer(ctx context.Context, cfg *dao.Answer) (string, error)
	GetExamById(ctx context.Context, examId string) (*dao.Exam, error)
	GetQuestionsByFilters(ctx context.Context, questionIds []string) ([]*dao.Question, error)
}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}

func (s *Service) ValidateSubmissionTime(ctx context.Context, startTime int64, endTime int64) (bool, error) {
	// Get the current time
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	// Calculate the end time plus 5 minutes
	endTimePlus5 := endTime + (5 * 60 * 1000) //5 minutes in milliseconds

	if currentTime < startTime {
		return false, errors.New("submission is not allowed before the exam start time")
	}

	if currentTime > endTimePlus5 {
		return false, errors.New("submission is not allowed after the exam end time")
	}

	if currentTime >= endTime && currentTime <= endTimePlus5 {
		// Submission is allowed
		return true, nil
	}
	// Submission is not allowed
	return false, nil
}

func (s *Service) CheckAnswers(ctx context.Context, questions []string, answers []dto.QuestionAnswer) (int, error) {
	//GetQuestions
	questionsList, err := s.repo.GetQuestionsByFilters(ctx, questions)
	if err != nil {
		return 0, err
	}

	// Create a map of question IDs to correct answers
	correctAnswers := make(map[string]string)
	for _, q := range questionsList {
		correctAnswers[q.ID] = q.Correct
	}

	//validate each answer
	correctCount := 0
	for _, answer := range answers {
		correctAnswer, ok := correctAnswers[answer.QuestionId]
		if !ok {
			logger.Error(ctx, "question with ID %s not found", answer.QuestionId)
			continue
		}
		if answer.Answer == correctAnswer {
			correctCount += 1
		}
	}
	return correctCount, nil
}

func (s *Service) CreateOrUpdateAnswer(ctx context.Context, req *dto.AnswerRequest) (*dto.AnswerResponse, error) {
	//validate examId
	exam, err := s.repo.GetExamById(ctx, req.ExamID)
	if err != nil {
		logger.Error(ctx, "failed to validate exam ID: %v", err)
		return nil, err
	}

	//validate time of submission
	valid, err := s.ValidateSubmissionTime(ctx, exam.StartTime, exam.EndTime)
	if err != nil {
		logger.Error(ctx, "Error validating submission: %v", err)
		return nil, err
	}
	if !valid {
		return nil, errors.New("submission is not allowed")
	}

	//Check Question answers
	corrected, err := s.CheckAnswers(ctx, exam.Questions, req.Answers)
	if err != nil {
		return nil, err
	}
	ans := req.ToMongoObject()

	ans.Result.Attempted = int64(len(req.Answers))
	ans.Result.Correct = int64(corrected)

	objectId, err := s.repo.CreateOrUpdateAnswer(ctx, ans)
	if err != nil {
		return nil, err
	}
	return &dto.AnswerResponse{StatusCode: http.StatusCreated, Id: objectId}, nil
}
