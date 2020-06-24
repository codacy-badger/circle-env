package usecases

import (
	"errors"
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/kou-pg-0131/circle-env/src/utils"
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
	cfg, err := u.configRepository.Get()
	if err != nil {
		return err
	}

	es, err := u.envsRepository.Load()
	if err != nil {
		return err
	}

	curs, err := u.envsRepository.All(cfg)
	if err != nil {
		return err
	}

	ds := curs.Compare(es)
	for _, d := range *ds {
		switch d.Status {
		case domain.NotChanged:
			fmt.Printf("%s = \"%s\"\n", d.Name, d.Before.Value)
		case domain.Changed:
			fmt.Print(utils.Colorf(
				"~ %s = \"%s\" -> \"%s\"\n",
				d.Name,
				d.Before.Value,
				d.After.Value,
			).Green().Bold().String())
		case domain.Added:
			fmt.Printf(utils.Colorf(
				"+ %s = \"%s\"\n",
				d.Name,
				d.After.Value,
			).Green().Bold().String())
		}
	}

	fmt.Println("")
	yes, err := utils.Confirm(utils.Color("Continue?(yes/no): ").Bold().String())
	if err != nil {
		return err
	}
	if !yes {
		return errors.New("cancelled")
	}

	for _, e := range *es {
		if curs.Has(e.Name) {
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
