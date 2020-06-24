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

func (c *CircleCIAPIClient) handleRequestError(res *HttpResponse, cfg *domain.Config) error {
	switch res.StatusCode {
	case 200, 201:
		return nil
	case 403:
		return errors.New("permission denied")
	case 404:
		return fmt.Errorf(`project not found: "%s"`, cfg.Slug())
	default:
		return fmt.Errorf("api request failed with status code %d: %s", res.StatusCode, string(res.Body))
	}
}

func (c *CircleCIAPIClient) GetEnvs(cfg *domain.Config) (*domain.Envs, error) {
	url := fmt.Sprintf(
		"https://circleci.com/api/v1.1/project/%s/envvar?circle-token=%s",
		cfg.Slug(),
		cfg.Token,
	)

	res, err := c.httpClient.Get(url, map[string]string{"Accept": "application/json"})
	if err != nil {
		return nil, err
	}

	if err := c.handleRequestError(res, cfg); err != nil {
		return nil, err
	}

	es := new(domain.Envs)
	if err := json.Unmarshal(res.Body, es); err != nil {
		return nil, err
	}

	return es, nil
}

func (c *CircleCIAPIClient) CreateEnv(cfg *domain.Config, e *domain.Env) error {
	url := fmt.Sprintf(
		"https://circleci.com/api/v1.1/project/%s/envvar?circle-token=%s",
		cfg.Slug(),
		cfg.Token,
	)

	bs, err := json.Marshal(e)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Post(url, map[string]string{"Content-Type": "application/json", "Accept": "application/json"}, bs)
	if err != nil {
		return err
	}

	if err := c.handleRequestError(res, cfg); err != nil {
		return err
	}

	return nil
}
