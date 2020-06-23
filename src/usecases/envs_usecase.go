package usecases

import "github.com/kou-pg-0131/circle-env/src/domain"

type IEnvsUsecase interface {
	ShowAll() (*domain.Envs, error)
}

type EnvsUsecase struct {
	envsRepository   IEnvsRepository
	configRepository IConfigRepository
}

type EnvsUsecaseOption struct {
	EnvsRepository   IEnvsRepository
	ConfigRepository IConfigRepository
}

func NewEnvsUsecase(opt *EnvsUsecaseOption) *EnvsUsecase {
	return &EnvsUsecase{
		envsRepository:   opt.EnvsRepository,
		configRepository: opt.ConfigRepository,
	}
}

func (u *EnvsUsecase) ShowAll() (*domain.Envs, error) {
	cfg, err := u.configRepository.Get()
	if err != nil {
		return nil, err
	}

	es, err := u.envsRepository.All(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}
