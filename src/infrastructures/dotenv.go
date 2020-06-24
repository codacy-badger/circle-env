package infrastructures

import (
	"github.com/joho/godotenv"
	"github.com/kou-pg-0131/circle-env/src/domain"
)

type Dotenv struct{}

func NewDotenv() *Dotenv {
	return new(Dotenv)
}

func (d *Dotenv) Load(path string) (*domain.Envs, error) {
	m, err := godotenv.Read(path)
	if err != nil {
		return nil, err
	}

	es := new(domain.Envs)
	for k, v := range m {
		*es = append(*es, &domain.Env{Name: k, Value: v})
	}

	return es, nil
}
