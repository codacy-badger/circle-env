package usecases

import "github.com/kou-pg-0131/circle-env/src/domain"

type IEnvsRepository interface {
	All(cfg *domain.Config) (*domain.Envs, error)
}
