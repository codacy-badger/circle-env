package infrastructures

import (
	"errors"
	"testing"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewCircleCIAPIClient()
 */

func Test_NewCircleCIAPIClient(t *testing.T) {
	c := NewCircleCIAPIClient()
	assert.NotNil(t, c.httpClient)
}

/*
 * CircleCIAPIClient.GetEnvs()
 */

func TestCircleCIAPIClient_GetEnvs_ReturnEnvs(t *testing.T) {
	expected := &domain.Envs{
		{Name: "NAME", Value: "VALUE"},
	}

	h := new(mockHTTPClient)
	h.On("Get",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Accept": "application/json"},
	).Return(&HTTPResponse{Body: []byte(`[{"name":"NAME","value":"VALUE"}]`), StatusCode: 200}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	es, err := c.GetEnvs(cfg)
	assert.Equal(t, expected, es)
	assert.Nil(t, err)
	h.AssertNumberOfCalls(t, "Get", 1)
}

func TestCircleCIAPIClient_GetEnvs_ReturnErrorWhenHTTPGetFailed(t *testing.T) {
	h := new(mockHTTPClient)
	h.On("Get",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Accept": "application/json"},
	).Return((*HTTPResponse)(nil), errors.New("SOMETHING_WRONG"))

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	es, err := c.GetEnvs(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	assert.Nil(t, es)
	h.AssertNumberOfCalls(t, "Get", 1)
}

func TestCircleCIAPIClient_GetEnvs_ReturnErrorWhen403Status(t *testing.T) {
	h := new(mockHTTPClient)
	h.On("Get",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Accept": "application/json"},
	).Return(&HTTPResponse{Body: []byte("BODY"), StatusCode: 403}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	es, err := c.GetEnvs(cfg)
	assert.Equal(t, errors.New("permission denied"), err)
	assert.Nil(t, es)
	h.AssertNumberOfCalls(t, "Get", 1)
}

func TestCircleCIAPIClient_GetEnvs_ReturnErrorWhen404Status(t *testing.T) {
	h := new(mockHTTPClient)
	h.On("Get",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Accept": "application/json"},
	).Return(&HTTPResponse{Body: []byte("BODY"), StatusCode: 404}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	es, err := c.GetEnvs(cfg)
	assert.EqualError(t, err, `project not found: "github/user/repo"`)
	assert.Nil(t, es)
	h.AssertNumberOfCalls(t, "Get", 1)
}

func TestCircleCIAPIClient_GetEnvs_ReturnErrorWhenUnexpectedStatus(t *testing.T) {
	h := new(mockHTTPClient)
	h.On("Get",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Accept": "application/json"},
	).Return(&HTTPResponse{Body: []byte("BODY"), StatusCode: 500}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	es, err := c.GetEnvs(cfg)
	assert.EqualError(t, err, "api request failed with status code 500: BODY")
	assert.Nil(t, es)
	h.AssertNumberOfCalls(t, "Get", 1)
}

/*
 * CircleCIAPIClient.CreateEnv()
 */

func TestCircleCIAPIClient_CreateEnv_ReturnNil(t *testing.T) {
	e := &domain.Env{
		Name: "NAME", Value: "VALUE",
	}

	h := new(mockHTTPClient)
	h.On("Post",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
		[]byte(`{"name":"NAME","value":"VALUE"}`),
	).Return(&HTTPResponse{StatusCode: 201, Body: []byte("BODY")}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	err := c.CreateEnv(cfg, e)
	assert.Nil(t, err)
	h.AssertNumberOfCalls(t, "Post", 1)
}

func TestCircleCIAPIClient_CreateEnv_ReturnErrorWhenRequestFailed(t *testing.T) {
	e := &domain.Env{
		Name: "NAME", Value: "VALUE",
	}

	h := new(mockHTTPClient)
	h.On("Post",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
		[]byte(`{"name":"NAME","value":"VALUE"}`),
	).Return((*HTTPResponse)(nil), errors.New("SOMETHING_WRONG"))

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	err := c.CreateEnv(cfg, e)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	h.AssertNumberOfCalls(t, "Post", 1)
}

func TestCircleCIAPIClient_CreateEnv_ReturnErrorWhen403Status(t *testing.T) {
	e := &domain.Env{
		Name: "NAME", Value: "VALUE",
	}

	h := new(mockHTTPClient)
	h.On("Post",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
		[]byte(`{"name":"NAME","value":"VALUE"}`),
	).Return(&HTTPResponse{StatusCode: 403, Body: []byte("BODY")}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	err := c.CreateEnv(cfg, e)
	assert.EqualError(t, err, "permission denied")
	h.AssertNumberOfCalls(t, "Post", 1)
}

func TestCircleCIAPIClient_CreateEnv_ReturnErrorWhen404Status(t *testing.T) {
	e := &domain.Env{
		Name: "NAME", Value: "VALUE",
	}

	h := new(mockHTTPClient)
	h.On("Post",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
		[]byte(`{"name":"NAME","value":"VALUE"}`),
	).Return(&HTTPResponse{StatusCode: 404, Body: []byte("BODY")}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	err := c.CreateEnv(cfg, e)
	assert.EqualError(t, err, `project not found: "github/user/repo"`)
	h.AssertNumberOfCalls(t, "Post", 1)
}

func TestCircleCIAPIClient_CreateEnv_ReturnErrorWhenUnexpectedStatus(t *testing.T) {
	e := &domain.Env{
		Name: "NAME", Value: "VALUE",
	}

	h := new(mockHTTPClient)
	h.On("Post",
		"https://circleci.com/api/v1.1/project/github/user/repo/envvar?circle-token=token",
		map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
		[]byte(`{"name":"NAME","value":"VALUE"}`),
	).Return(&HTTPResponse{StatusCode: 500, Body: []byte("BODY")}, nil)

	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}
	c := &CircleCIAPIClient{httpClient: h}

	err := c.CreateEnv(cfg, e)
	assert.EqualError(t, err, "api request failed with status code 500: BODY")
	h.AssertNumberOfCalls(t, "Post", 1)
}
