package usecases

import (
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

type IEnvsUsecase interface {
	ShowAll() (*domain.Envs, error)
	Push() error
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

func (u *EnvsUsecase) Push() error {
	es, err := u.envsRepository.Load()
	if err != nil {
		return err
	}

	cfg, err := u.configRepository.Get()
	if err != nil {
		return err
	}

	curs, err := u.envsRepository.All(cfg)
	if err != nil {
		return err
	}

	for _, e := range *es {
		if curs.Has(e) {
			fmt.Printf("Saving `%s`...\n", e.Name)
		} else {
			fmt.Printf("Creating `%s`...\n", e.Name)
		}

		if err := u.envsRepository.Save(cfg, e); err != nil {
			return err
		}
	}

	return nil
}
