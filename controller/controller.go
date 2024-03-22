package controller

import (
	"context"
	"examservice/controller/http"
)

type Controller interface {
	Listen(ctx context.Context) error
}

type Config struct {
	Name string
	HTTP http.Config
}

func NewController(ctx context.Context, conf *Config) Controller {
	return nil
}
