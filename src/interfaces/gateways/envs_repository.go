package gateways

import (
	"errors"
	"os"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

const (
	DotenvPath string = ".circle-env/.env"
)

type EnvsRepository struct {
	dotenv    IDotenv
	apiClient IAPIClient
}

func NewEnvsRepository(c IAPIClient, d IDotenv) *EnvsRepository {
	return &EnvsRepository{apiClient: c, dotenv: d}
}

func (r *EnvsRepository) All(cfg *domain.Config) (*domain.Envs, error) {
	es, err := r.apiClient.GetEnvs(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (r *EnvsRepository) Load() (*domain.Envs, error) {
	_, err := os.Stat(DotenvPath)
	if err != nil {
		return nil, errors.New("`.circle-env/.env` not found")
	}

	es, err := r.dotenv.Load(DotenvPath)
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (r *EnvsRepository) Save(cfg *domain.Config, e *domain.Env) error {
	return r.apiClient.CreateEnv(cfg, e)
}

func (r *EnvsRepository) Delete(cfg *domain.Config, name string) error {
	return r.apiClient.DeleteEnv(cfg, name)
}
