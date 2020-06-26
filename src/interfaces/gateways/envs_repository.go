package gateways

import (
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

const (
	// DotenvPath ...
	DotenvPath string = ".circle-env/.env"
)

// EnvsRepository ...
type EnvsRepository struct {
	fs        IFileSystem
	dotenv    IDotenv
	apiClient IAPIClient
}

// NewEnvsRepository ...
func NewEnvsRepository(c IAPIClient, fs IFileSystem, d IDotenv) *EnvsRepository {
	return &EnvsRepository{apiClient: c, fs: fs, dotenv: d}
}

// All ...
func (r *EnvsRepository) All(cfg *domain.Config) (*domain.Envs, error) {
	es, err := r.apiClient.GetEnvs(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}

// Load ...
func (r *EnvsRepository) Load() (*domain.Envs, error) {
	if !r.fs.IsExists(DotenvPath) {
		return nil, fmt.Errorf("`%s` not found", DotenvPath)
	}

	dot, err := r.dotenv.Load(DotenvPath)
	if err != nil {
		return nil, err
	}

	es := new(domain.Envs)
	for _, e := range *dot {
		if e.Value != "" {
			*es = append(*es, e)
		}
	}

	return es, nil
}

// Save ...
func (r *EnvsRepository) Save(cfg *domain.Config, e *domain.Env) error {
	return r.apiClient.CreateEnv(cfg, e)
}

// Delete ...
func (r *EnvsRepository) Delete(cfg *domain.Config, name string) error {
	return r.apiClient.DeleteEnv(cfg, name)
}
