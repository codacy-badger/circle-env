package presenters

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

type IEnvsPresenter interface {
	Print(es *domain.Envs) error
}

type PlainEnvsPresenter struct{}

func NewEnvsPresenter(j bool) IEnvsPresenter {
	if j {
		return new(JSONEnvsPresenter)
	} else {
		return new(PlainEnvsPresenter)
	}
}

func (p *PlainEnvsPresenter) Print(es *domain.Envs) error {
	if len(*es) == 0 {
		fmt.Println("No environment variables are set.")
		return nil
	}

	for _, e := range *es {
		fmt.Printf("%s = \"%s\"\n", e.Name, e.Value)
	}

	return nil
}

type JSONEnvsPresenter struct{}

func (p *JSONEnvsPresenter) Print(es *domain.Envs) error {
	bs, err := json.Marshal(es)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = json.Indent(buf, bs, "", "  "); err != nil {
		return err
	}

	fmt.Println(buf.String())

	return nil
}
