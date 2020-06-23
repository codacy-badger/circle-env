package gateways

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

type ConfigRepository struct{}

func NewConfigRepository() *ConfigRepository {
	return &ConfigRepository{}
}

func (r *ConfigRepository) Save(c *domain.Config) error {
	if err := os.MkdirAll(".circle-env", os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(".circle-env/config")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(c.Ini()))
	if err != nil {
		return err
	}

	fmt.Println("created `.circle-env/config`.")
	return nil
}

func (r *ConfigRepository) Get() (*domain.Config, error) {
	i, err := ini.Load(".circle-env/config")
	if err != nil {
		return nil, err
	}

	vcs := (domain.VCS)(i.Section("").Key("vcs").String())
	if !vcs.IsValid() {
		return nil, domain.ErrInvalidVCS
	}

	return &domain.Config{
		VCS:   vcs,
		Token: i.Section("").Key("token").String(),
		Repo:  i.Section("").Key("repo").String(),
		User:  i.Section("").Key("user").String(),
	}, nil
}
