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
		"https://circleci.com/api/v1.1/project/%s/envvar?circle-token=%s",
		cfg.Slug(),
		cfg.Token,
	)

	res, err := c.httpClient.Get(url, map[string]string{"Accept": "application/json"})
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("api request failed with status code %d\nBody: %s", res.StatusCode, string(res.Body))
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

	if res.StatusCode != 201 {
		return fmt.Errorf("api request failed with status code %d\nBody: %s", res.StatusCode, string(res.Body))
	}

	return nil
}

func (c *CircleCIAPIClient) DeleteEnv(cfg *domain.Config, name string) error {
	url := fmt.Sprintf(
		"https://circleci.com/api/v1.1/project/%s/envvar/%s?circle-token=%s",
		cfg.Slug(),
		name,
		cfg.Token,
	)

	res, err := c.httpClient.Delete(url, map[string]string{"Accept": "application/json"})
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("api request failed with status code %d\nBody: %s", res.StatusCode, string(res.Body))
	}

	return nil
}
