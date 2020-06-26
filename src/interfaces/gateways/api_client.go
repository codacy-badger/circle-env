package gateways

import (
	"github.com/kou-pg-0131/circle-env/src/domain"
)

// IAPIClient ...
type IAPIClient interface {
	GetEnvs(cfg *domain.Config) (*domain.Envs, error)
	CreateEnv(cfg *domain.Config, e *domain.Env) error
	DeleteEnv(cfg *domain.Config, name string) error
}
