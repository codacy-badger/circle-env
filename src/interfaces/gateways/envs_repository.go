package gateways

import (
	"github.com/kou-pg-0131/circle-env/src/domain"
)

type EnvsRepository struct {
	apiClient IAPIClient
}

func NewEnvsRepository(c IAPIClient) *EnvsRepository {
	return &EnvsRepository{apiClient: c}
}

func (r *EnvsRepository) All(cfg *domain.Config) (*domain.Envs, error) {
	es, err := r.apiClient.GetEnvs(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}
