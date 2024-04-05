package controller

import (
	"context"
	"examservice/controller/http"
	"examservice/service"

	"github.com/bappaapp/goutils/logger"
)

const (
	HTTP = "HTTP"
)

type Controller interface {
	Listen(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type Config struct {
	Name string
	HTTP http.Config
}

func NewController(ctx context.Context, cfg *Config, srvFactory *service.ServiceFactory) Controller {
	switch cfg.Name {
	case HTTP:
		return http.NewHTTPController(ctx, &cfg.HTTP, srvFactory)
	default:
		logger.Panic(ctx, "invalid controller name")
	}
	return nil
}
