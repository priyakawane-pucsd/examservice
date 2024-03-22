package examservice

import "context"

type Service struct {
	conf *Config
	repo Repository
}

type Config struct{}

type Repository interface{}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}
