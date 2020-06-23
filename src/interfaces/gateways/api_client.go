package gateways

import (
	"github.com/kou-pg-0131/circle-env/src/domain"
)

type IAPIClient interface {
	GetEnvs(c *domain.Config) (*domain.Envs, error)
}
