package mongo

import "context"

type Repository struct {
}

type Config struct {
}

func NewRepository(ctx context.Context, conf *Config) *Repository {
	return &Repository{}
}
