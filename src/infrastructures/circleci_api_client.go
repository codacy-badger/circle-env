package infrastructures

import (
	"encoding/json"
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

type CircleCIAPIClient struct {
	httpClient IHttpClient
}

func NewCircleCIAPIClient() *CircleCIAPIClient {
	return &CircleCIAPIClient{
		httpClient: NewHttpClient(),
	}
}

func (c *CircleCIAPIClient) GetEnvs(cfg *domain.Config) (*domain.Envs, error) {
	url := fmt.Sprintf(
		"https://circleci.com/api/v1.1/project/%s/%s/%s/envvar?circle-token=%s",
		cfg.VCS,
		cfg.User,
		cfg.Repo,
		cfg.Token,
	)

	res, err := c.httpClient.Get(url, map[string]string{"Accept": "application/json"})
	if err != nil {
		return nil, err
	}

	es := new(domain.Envs)
	if err := json.Unmarshal(res.Body, es); err != nil {
		return nil, err
	}

	return es, nil
}
