package controllers

import (
	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
	"github.com/kou-pg-0131/circle-env/src/usecases"
	"github.com/kou-pg-0131/circle-env/src/utils"
)

type ConfigController struct {
	usecase usecases.IConfigUsecase
	scanner utils.IScanner
}

func NewConfigController() *ConfigController {
	return &ConfigController{
		scanner: utils.NewScanner(),
		usecase: usecases.NewConfigUsecase(
			gateways.NewConfigRepository(),
		),
	}
}

func (c *ConfigController) Initialize() {
	err := c.usecase.Initialize()
	if err != nil {
		panic(err)
	}
}
