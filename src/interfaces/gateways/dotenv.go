package gateways

import "github.com/kou-pg-0131/circle-env/src/domain"

// IDotenv ...
type IDotenv interface {
	Load(path string) (*domain.Envs, error)
}
