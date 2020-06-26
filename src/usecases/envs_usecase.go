package usecases

import (
	"errors"
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/kou-pg-0131/circle-env/src/utils"
)

type IEnvsUsecase interface {
	ShowAll() (*domain.Envs, error)
	Sync(del bool) error
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

func (u *EnvsUsecase) Sync(del bool) error {
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

	ds := curs.Compare(es, del)
	for _, d := range *ds {
		switch d.Status {
		case domain.NotChanged:
			fmt.Printf("%s = \"%s\"\n", d.Name, d.Before)
		case domain.Deleted:
			fmt.Printf(
				"%s%s\n",
				utils.Colorf("- %s = \"%s\" -> ", d.Name, d.Before).Red().Bold().String(),
				utils.Colorf("null").Secondary().String(),
			)
		case domain.Changed:
			fmt.Println(utils.Colorf(
				"~ %s = \"%s\" -> \"%s\"",
				d.Name,
				d.Before,
				d.After,
			).Green().Bold().String())
		case domain.Added:
			fmt.Println(utils.Colorf(
				"+ %s = \"%s\"",
				d.Name,
				d.After,
			).Green().Bold().String())
		}
	}

	fmt.Println("")
	yes, err := utils.Confirm(utils.Colorf("Continue?(yes/no): ").Bold().String())
	if err != nil {
		return err
	}
	if !yes {
		return errors.New("cancelled")
	}

	for _, d := range *ds {
		switch d.Status {
		case domain.Added:
			fmt.Printf("Creating `%s`...\n", d.Name)
			if err := u.envsRepository.Save(cfg, &domain.Env{Name: d.Name, Value: d.After}); err != nil {
				return err
			}
		case domain.Changed:
			fmt.Printf("Modifying `%s`...\n", d.Name)
			if err := u.envsRepository.Save(cfg, &domain.Env{Name: d.Name, Value: d.After}); err != nil {
				return err
			}
		case domain.Deleted:
			fmt.Printf("Deleting `%s`...\n", d.Name)
			if err := u.envsRepository.Delete(cfg, d.Name); err != nil {
				return err
			}
		}
	}

	fmt.Println(utils.Colorf("Completed!").Bold().String())

	return nil
}
