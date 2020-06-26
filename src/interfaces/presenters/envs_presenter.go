package presenters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

// IEnvsPresenter ...
type IEnvsPresenter interface {
	String(es *domain.Envs) (string, error)
}

// PlainEnvsPresenter ...
type PlainEnvsPresenter struct{}

// NewEnvsPresenter ...
func NewEnvsPresenter(j bool) IEnvsPresenter {
	if j {
		return new(JSONEnvsPresenter)
	}

	return new(PlainEnvsPresenter)
}

func (p *PlainEnvsPresenter) String(es *domain.Envs) (string, error) {
	if len(*es) == 0 {
		return "no environment variables are set.", nil
	}

	ss := []string{}
	for _, e := range *es {
		ss = append(ss, fmt.Sprintf("%s = \"%s\"", e.Name, e.Value))
	}

	return strings.Join(ss, "\n"), nil
}

// JSONEnvsPresenter ...
type JSONEnvsPresenter struct{}

func (p *JSONEnvsPresenter) String(es *domain.Envs) (string, error) {
	bs, err := json.Marshal(es)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = json.Indent(buf, bs, "", "  "); err != nil {
		return "", err
	}

	return buf.String(), nil
}
