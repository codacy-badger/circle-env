package controllers

import (
	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
	"github.com/kou-pg-0131/circle-env/src/usecases"
)

type ConfigController struct {
	usecase usecases.IConfigUsecase
}

func NewConfigController(fs gateways.IFileSystem) *ConfigController {
	return &ConfigController{
		usecase: usecases.NewConfigUsecase(
			gateways.NewConfigRepository(fs),
		),
	}
}

func (c *ConfigController) Initialize() error {
	return c.usecase.Initialize()
}
