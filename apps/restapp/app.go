package restapp

import (
	"context"
	"examservice/config"
	"examservice/controller"
	"examservice/repository"
	"examservice/service"

	"github.com/bappaapp/goutils/configutils"
	"github.com/bappaapp/goutils/logger"
)

type RestApp struct {
	ctrl controller.Controller
}

func NewRestApp(ctx context.Context, configpath string) *RestApp {
	var conf config.Config
	err := configutils.ReadConfig(configpath, &conf)
	if err != nil {
		logger.Panic(ctx, "failed to load config: %v", err.Error())
	}
	repo := repository.NewRepository(ctx, &conf.Repository)
	services := service.NewServiceFactory(ctx, &conf.Service, repo)
	controller := controller.NewController(ctx, &conf.Controller, services)

	return &RestApp{ctrl: controller}
}

func (app *RestApp) Start(ctx context.Context) error {
	return app.ctrl.Listen(ctx)
}

func (app *RestApp) Shutdown(ctx context.Context) error {
	return app.ctrl.Shutdown(ctx)
}
