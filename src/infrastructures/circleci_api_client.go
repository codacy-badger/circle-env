package infrastructures

import (
	"encoding/json"
	"errors"
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

	switch res.StatusCode {
	case 200:
	case 403:
		return nil, errors.New("permission denied")
	case 404:
		return nil, errors.New("project not found")
	case 429:
		return nil, errors.New("too many requests")
	default:
		return nil, fmt.Errorf("request failed with status code %d\nBody: %s", res.StatusCode, string(res.Body))
	}

	es := new(domain.Envs)
	if err := json.Unmarshal(res.Body, es); err != nil {
		return nil, err
	}

	return es, nil
}
