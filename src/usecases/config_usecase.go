package usecases

import (
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/kou-pg-0131/circle-env/src/utils"
)

type IConfigUsecase interface {
	Initialize() error
}

type ConfigUsecase struct {
	scanner    utils.IScanner
	repository IConfigRepository
}

func NewConfigUsecase(r IConfigRepository) *ConfigUsecase {
	return &ConfigUsecase{
		repository: r,
		scanner:    utils.NewScanner(),
	}
}

func (u *ConfigUsecase) Initialize() error {
	var vcs domain.VCS

	for {
		fmt.Print("VCS (`github` or `bitbucket`): ")
		s, err := u.scanner.Scan()
		if err != nil {
			return err
		}

		if (domain.VCS)(s).IsValid() {
			vcs = (domain.VCS)(s)
			break
		} else {
			fmt.Println("`github` or `bitbucket`")
		}
	}

	fmt.Print("User: ")
	user, err := u.scanner.Scan()
	if err != nil {
		return err
	}

	fmt.Print("Repository: ")
	repo, err := u.scanner.Scan()
	if err != nil {
		return err
	}

	fmt.Print("Token: ")
	tkn, err := u.scanner.Scan()
	if err != nil {
		return err
	}

	return u.repository.Save(&domain.Config{
		VCS:   vcs,
		User:  user,
		Repo:  repo,
		Token: tkn,
	})
}
