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
	repository IConfigRepository
}

func NewConfigUsecase(r IConfigRepository) *ConfigUsecase {
	return &ConfigUsecase{
		repository: r,
	}
}

func (u *ConfigUsecase) Initialize() error {
	var vcs domain.VCS

	for {
		fmt.Print("VCS (`github` or `bitbucket`): ")
		s, err := utils.Scanner.Scan()
		if err != nil {
			return err
		}

		if (domain.VCS)(s).IsValid() {
			vcs = (domain.VCS)(s)
			break
		}

		fmt.Println("Please enter `github` or `bitbucket`.")
	}

	fmt.Print("User: ")
	user, err := utils.Scanner.Scan()
	if err != nil {
		return err
	}

	fmt.Print("Repository: ")
	repo, err := utils.Scanner.Scan()
	if err != nil {
		return err
	}

	fmt.Print("CircleCI API Token: ")
	tkn, err := utils.Scanner.Scan()
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
