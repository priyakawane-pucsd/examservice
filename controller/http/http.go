package http

import (
	"context"
	"examservice/controller/http/ping"
	"examservice/controller/http/swagger"
	"examservice/service"
	"fmt"
	"log"
	"net/http"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port         int
	GinModeDebug bool
}

type HTTPController struct {
	conf       *Config
	srvFactory *service.ServiceFactory
	server     *http.Server
}

func NewHTTPController(ctx context.Context, conf *Config, srvFactory *service.ServiceFactory) *HTTPController {
	return &HTTPController{conf: conf, srvFactory: srvFactory}
}

func (c *HTTPController) Listen(ctx context.Context) error {
	if !c.conf.GinModeDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	//registering controllers
	ping.NewPingController(ctx, c.srvFactory.GetPingService()).Register(router)
	swagger.NewSwaggerController(ctx).Register(router)

	logger.Info(ctx, "swagger link: http://localhost:%d/examservice/swagger/index.html", c.conf.Port)
	c.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", c.conf.Port),
		Handler: router,
	}
	log.Printf("HTTP server started listening on :%d", c.conf.Port)
	err := c.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Panic(ctx, "failed to start the server, Err: %v", err.Error())
		return err
	}
	return nil
}

func (c *HTTPController) Shutdown(ctx context.Context) error {
	return nil
}
