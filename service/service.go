package service

import (
	"context"
	"examservice/service/examservice"
	"examservice/service/questionservice"
)

type ServiceFactory struct {
	questionService *questionservice.Service
	examService     *examservice.Service
}

type Repository interface {
	examservice.Repository
	questionservice.Repository
}

type Config struct {
	QuestionService questionservice.Config
	ExamService     examservice.Config
}

func NewServiceFactory(ctx context.Context, conf *Config, repo Repository) *ServiceFactory {
	return &ServiceFactory{
		examService:     examservice.NewService(ctx, &conf.ExamService, repo),
		questionService: questionservice.NewService(ctx, &conf.QuestionService, repo),
	}
}
