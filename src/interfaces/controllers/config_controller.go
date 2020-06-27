package controllers

import (
	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
	"github.com/kou-pg-0131/circle-env/src/usecases"
)

// IConfigController ...
type IConfigController interface {
	Initialize() error
}

// ConfigController ...
type ConfigController struct {
	usecase usecases.IConfigUsecase
}

// NewConfigController ...
func NewConfigController(fs gateways.IFileSystem) *ConfigController {
	return &ConfigController{
		usecase: usecases.NewConfigUsecase(
			gateways.NewConfigRepository(fs),
		),
	}
}

// Initialize ...
func (c *ConfigController) Initialize() error {
	return c.usecase.Initialize()
}
