package gateways

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

const (
	// DirPath ...
	DirPath string = ".circle-env"
	// ConfigPath ...
	ConfigPath string = ".circle-env/config"
	// TokenPath ...
	TokenPath string = ".circle-env/circle-token"
)

// ConfigRepository ...
type ConfigRepository struct {
	fs IFileSystem
}

// NewConfigRepository ...
func NewConfigRepository(fs IFileSystem) *ConfigRepository {
	return &ConfigRepository{fs: fs}
}

// Save ...
func (r *ConfigRepository) Save(cfg *domain.Config) error {
	if err := r.fs.Mkdir(DirPath); err != nil {
		return err
	}

	c, err := r.fs.Create(ConfigPath)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.Write([]byte(cfg.Ini()))
	if err != nil {
		return err
	}

	t, err := r.fs.Create(TokenPath)
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

// Get ...
func (r *ConfigRepository) Get() (*domain.Config, error) {
	if !r.fs.IsExists(ConfigPath) {
		return nil, fmt.Errorf("`%s` not found, run `circle-env init`", ConfigPath)
	}

	if !r.fs.IsExists(TokenPath) {
		return nil, fmt.Errorf("`%s` not found, run `circle-env init`", TokenPath)
	}

	cbs, err := r.fs.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	i, err := ini.Load(cbs)
	if err != nil {
		return nil, err
	}

	v := i.Section("").Key("vcs").String()
	vcs, err := domain.VCSFromString(v)
	if err != nil {
		return nil, fmt.Errorf("`%s` is invalid vcs type, please check `%s`", v, ConfigPath)
	}

	tbs, err := r.fs.ReadFile(TokenPath)
	if err != nil {
		return nil, err
	}
	tkn := strings.TrimSpace(string(tbs))

	return &domain.Config{
		VCS:   vcs,
		User:  i.Section("").Key("user").String(),
		Repo:  i.Section("").Key("repo").String(),
		Token: tkn,
	}, nil
}
