package gateways

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/ini.v1"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

const (
	DirPath    string = ".circle-env"
	ConfigPath string = ".circle-env/config"
	TokenPath  string = ".circle-env/circle-token"
)

type ConfigRepository struct{}

func NewConfigRepository() *ConfigRepository {
	return &ConfigRepository{}
}

func (r *ConfigRepository) Save(cfg *domain.Config) error {
	if err := os.MkdirAll(DirPath, os.ModePerm); err != nil {
		return err
	}

	c, err := os.Create(ConfigPath)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Write([]byte(cfg.Ini()))
	if err != nil {
		return err
	}

	t, err := os.Create(TokenPath)
	if err != nil {
		return err
	}
	defer t.Close()

	_, err = t.Write([]byte(cfg.Token))
	if err != nil {
		return err
	}

	fmt.Printf("\ncreated config files.\n  - `%s`\n  - `%s`\n", ConfigPath, TokenPath)
	return nil
}

func (r *ConfigRepository) Get() (*domain.Config, error) {
	_, err := os.Stat(ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("`%s` not found, run `circle-env init`", ConfigPath)
	}

	_, err = os.Stat(TokenPath)
	if err != nil {
		return nil, fmt.Errorf("`%s` not found, run `circle-env init`", TokenPath)
	}

	i, err := ini.Load(ConfigPath)
	if err != nil {
		return nil, err
	}

	vcs := (domain.VCS)(i.Section("").Key("vcs").String())
	if !vcs.IsValid() {
		return nil, fmt.Errorf("`%s` is invalid vcs type, please check `.circle-env/config`", vcs)
	}

	bs, err := ioutil.ReadFile(TokenPath)
	if err != nil {
		return nil, err
	}
	tkn := string(bs)

	return &domain.Config{
		VCS:   vcs,
		Token: tkn,
		Repo:  i.Section("").Key("repo").String(),
		User:  i.Section("").Key("user").String(),
	}, nil
}
