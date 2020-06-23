package controllers

import (
	"github.com/kou-pg-0131/circle-env/src/interfaces/gateways"
	"github.com/kou-pg-0131/circle-env/src/interfaces/presenters"
	"github.com/kou-pg-0131/circle-env/src/usecases"
)

type EnvsController struct {
	usecase   usecases.IEnvsUsecase
	presenter presenters.IEnvsPresenter
}

func NewEnvsController(c gateways.IAPIClient) *EnvsController {
	return &EnvsController{
		presenter: presenters.NewEnvsPresenter(),
		usecase: usecases.NewEnvsUsecase(&usecases.EnvsUsecaseOption{
			EnvsRepository:   gateways.NewEnvsRepository(c),
			ConfigRepository: gateways.NewConfigRepository(),
		}),
	}
}

func (c *EnvsController) Show() {
	es, err := c.usecase.ShowAll()
	if err != nil {
		panic(err)
	}

	c.presenter.Print(es)
}
