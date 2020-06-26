package infrastructures

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kou-pg-0131/circle-env/src/domain"
)

// CircleCIAPIClient ...
type CircleCIAPIClient struct {
	httpClient IHTTPClient
}

// NewCircleCIAPIClient ...
func NewCircleCIAPIClient() *CircleCIAPIClient {
	return &CircleCIAPIClient{
		httpClient: NewHTTPClient(),
	}
}

func (c *CircleCIAPIClient) handleRequestError(res *HTTPResponse, cfg *domain.Config) error {
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

// GetEnvs ...
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

	es.Sort()
	return es, nil
}

// CreateEnv ...
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

// DeleteEnv ...
func (c *CircleCIAPIClient) DeleteEnv(cfg *domain.Config, name string) error {
	url := fmt.Sprintf(
		"https://circleci.com/api/v1.1/project/%s/envvar/%s?circle-token=%s",
		cfg.Slug(),
		name,
		cfg.Token,
	)

	res, err := c.httpClient.Delete(url, map[string]string{"Content-Type": "application/json", "Accept": "application/json"})
	if err != nil {
		return err
	}

	if err := c.handleRequestError(res, cfg); err != nil {
		return err
	}

	return nil
}
