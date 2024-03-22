package questionservice

import "context"

type Service struct {
	conf *Config
	repo Repository
}

type Repository interface {
}

type Config struct{}

func NewService(ctx context.Context, conf *Config, repo Repository) *Service {
	return &Service{conf: conf, repo: repo}
}
