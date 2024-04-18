package service

import (
	"context"
	"examservice/service/examservice"
	"examservice/service/pingservice"
	"examservice/service/questionservice"
)

type ServiceFactory struct {
	pingService     *pingservice.Service
	questionService *questionservice.Service
	examService     *examservice.Service
}

type Repository interface {
	pingservice.Repository
	examservice.Repository
	questionservice.Repository
}

type Config struct {
	QuestionService questionservice.Config
	ExamService     examservice.Config
}

func NewServiceFactory(ctx context.Context, conf *Config, repo Repository) *ServiceFactory {
	return &ServiceFactory{
		pingService:     pingservice.NewService(ctx, repo),
		examService:     examservice.NewService(ctx, &conf.ExamService, repo),
		questionService: questionservice.NewService(ctx, &conf.QuestionService, repo),
	}

}

func (sf *ServiceFactory) GetPingService() *pingservice.Service {
	return sf.pingService
}

func (sf *ServiceFactory) GetQuestionService() *questionservice.Service {
	return sf.questionService
}

func (sf *ServiceFactory) GetExamService() *examservice.Service {
	return sf.examService
}
