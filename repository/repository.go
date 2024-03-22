package repository

import (
	"context"
	"examservice/repository/mongo"
	"examservice/service"

	"github.com/bappaapp/goutils/logger"
)

const (
	MONGO = "Mongo"
)

type Config struct {
	Name  string
	Mongo mongo.Config
}

type Repository interface {
	service.Repository
}

func NewRepository(ctx context.Context, conf *Config) Repository {
	switch conf.Name {
	case MONGO:
		return mongo.NewRepository(ctx, &conf.Mongo)
	default:
		logger.Panic(ctx, "invalid repo name")
	}
	return nil
}
