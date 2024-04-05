package main

import (
	"context"
	"examservice/apps/restapp"
	"flag"

	"github.com/bappaapp/goutils/configutils"
	"github.com/bappaapp/goutils/logger"
)

type Application interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type Config struct {
	App string
}

var (
	configfile string
)

const (
	RESTAPP = "RESTAPP"
)

func main() {
	flag.StringVar(&configfile, "configfile", "dev.yml", "config file path")
	flag.Parse()
	var conf Config
	var ctx context.Context = context.Background()
	err := configutils.ReadConfig(configfile, &conf)
	if err != nil {
		logger.Panic(ctx, "failed to load config: %v", err.Error())
	}
	if conf.App == "" {
		conf.App = RESTAPP
	}
	var application Application
	switch conf.App {
	case RESTAPP:
		application = restapp.NewRestApp(ctx, configfile)
	default:
		logger.Panic(ctx, "invalid application name")
	}
	err = application.Start(ctx)
	if err != nil {
		logger.Panic(ctx, "failed to run application, Err: %v", err.Error())
	}
}
