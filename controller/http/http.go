package http

import "context"

type Config struct {
	Port int
}

type HTTPController struct {
	conf *Config
}

func NewHTTPController(ctx context.Context, conf *Config) *HTTPController {
	return &HTTPController{conf: conf}
}
