package presenters

import (
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

type IEnvsPresenter interface {
	Print(es *domain.Envs) error
}

type PlainEnvsPresenter struct{}

func NewEnvsPresenter() IEnvsPresenter {
	return new(PlainEnvsPresenter)
}

func (p *PlainEnvsPresenter) Print(es *domain.Envs) error {
	for _, e := range *es {
		fmt.Printf("%s = %s\n", e.Name, e.Value)
	}

	return nil
}
