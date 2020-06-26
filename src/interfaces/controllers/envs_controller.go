package controllers

import (
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
	"github.com/kou-pg-0131/circle-env/src/interfaces/presenters"
	"github.com/kou-pg-0131/circle-env/src/usecases"
)

// EnvsController ...
type EnvsController struct {
	usecase usecases.IEnvsUsecase
}

// NewEnvsController ...
func NewEnvsController(c gateways.IAPIClient, fs gateways.IFileSystem, d gateways.IDotenv) *EnvsController {
	return &EnvsController{
		usecase: usecases.NewEnvsUsecase(&usecases.EnvsUsecaseOption{
			EnvsRepository:   gateways.NewEnvsRepository(c, fs, d),
			ConfigRepository: gateways.NewConfigRepository(fs),
		}),
	}
}

// Show ...
func (c *EnvsController) Show(j bool) error {
	es, err := c.usecase.ShowAll()
	if err != nil {
		return err
	}

	s, err := presenters.NewEnvsPresenter(j).String(es)
	if err != nil {
		return err
	}

	fmt.Println(s)
	return nil
}

// Sync ...
func (c *EnvsController) Sync(del, noconf bool) error {
	return c.usecase.Sync(del, noconf)
}
