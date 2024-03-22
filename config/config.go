package config

import (
	"examservice/controller"
	"examservice/repository"
	"examservice/service"
)

type Config struct {
	Name       string
	Build      string
	App        string
	Controller controller.Config
	Service    service.Config
	Repository repository.Config
}
