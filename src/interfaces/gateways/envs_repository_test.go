package gateways

import (
	"errors"
	"testing"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewEnvsRepository()
 */

func Test_NewEnvsRepository(t *testing.T) {
	api := new(mockAPIClient)
	fs := new(mockFileSystem)
	d := new(mockDotenv)

	r := NewEnvsRepository(api, fs, d)
	assert.Equal(t, &EnvsRepository{
		apiClient: api,
		fs:        fs,
		dotenv:    d,
	}, r)
}

/*
 * EnvsRepository.All()
 */

func TestEnvsRepository_All_ReturnEnvs(t *testing.T) {
	expected := new(domain.Envs)
	cfg := new(domain.Config)

	api := new(mockAPIClient)
	api.On("GetEnvs", cfg).Return(expected, nil)

	r := &EnvsRepository{apiClient: api}

	es, err := r.All(cfg)
	assert.Equal(t, expected, es)
	assert.Nil(t, err)
	api.AssertNumberOfCalls(t, "GetEnvs", 1)
}

func TestEnvsRepository_All_ReturnErrorWhenGetEnvsFailed(t *testing.T) {
	cfg := new(domain.Config)

	api := new(mockAPIClient)
	api.On("GetEnvs", cfg).Return((*domain.Envs)(nil), errors.New("SOMETHING_WRONG"))

	r := &EnvsRepository{apiClient: api}

	es, err := r.All(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	assert.Nil(t, es)
	api.AssertNumberOfCalls(t, "GetEnvs", 1)
}

/*
 * EnvsRepository.Load()
 */

func TestEnvsRepository_Load_ReturnEnvs(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/.env").Return(true)

	de := new(mockDotenv)
	de.On("Load", ".circle-env/.env").Return(&domain.Envs{
		{Name: "HOGE", Value: "PIYO"},
		{Name: "FOO", Value: "BAR"},
		{Name: "EMPTY", Value: ""},
	}, nil)

	r := &EnvsRepository{fs: fs, dotenv: de}

	es, err := r.Load()
	assert.Equal(t, &domain.Envs{
		{Name: "HOGE", Value: "PIYO"},
		{Name: "FOO", Value: "BAR"},
	}, es)
	assert.Nil(t, err)
	fs.AssertNumberOfCalls(t, "IsExists", 1)
	de.AssertNumberOfCalls(t, "Load", 1)
}

func TestEnvsRepository_Load_ReturnErrorWhenDotenvNotFound(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/.env").Return(false)

	de := new(mockDotenv)

	r := &EnvsRepository{fs: fs, dotenv: de}

	es, err := r.Load()
	assert.Nil(t, es)
	assert.EqualError(t, err, "`.circle-env/.env` not found")
	fs.AssertNumberOfCalls(t, "IsExists", 1)
	de.AssertNumberOfCalls(t, "Load", 0)
}

func TestEnvsRepository_Load_ReturnErrorWhenLoadDotenvFailed(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/.env").Return(true)

	de := new(mockDotenv)
	de.On("Load", ".circle-env/.env").Return((*domain.Envs)(nil), errors.New("SOMETHING_WRONG"))

	r := &EnvsRepository{fs: fs, dotenv: de}

	es, err := r.Load()
	assert.Nil(t, es)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	fs.AssertNumberOfCalls(t, "IsExists", 1)
	de.AssertNumberOfCalls(t, "Load", 1)
}

/*
 * EnvsRepository.Save()
 */

func TestEnvsRepository_Save_ReturnNil(t *testing.T) {
	cfg := new(domain.Config)
	e := new(domain.Env)
	api := new(mockAPIClient)
	api.On("CreateEnv", cfg, e).Return(nil)

	r := &EnvsRepository{apiClient: api}

	err := r.Save(cfg, e)
	assert.Nil(t, err)
	api.AssertNumberOfCalls(t, "CreateEnv", 1)
}

func TestEnvsRepository_Save_ReturnErrorWhenCreateEnvFailed(t *testing.T) {
	cfg := new(domain.Config)
	e := new(domain.Env)
	api := new(mockAPIClient)
	api.On("CreateEnv", cfg, e).Return(errors.New("SOMETHING_WRONG"))

	r := &EnvsRepository{apiClient: api}

	err := r.Save(cfg, e)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	api.AssertNumberOfCalls(t, "CreateEnv", 1)
}

/*
 * EnvsRepository.Delete()
 */

func TestEnvsRepository_Delete_ReturnNil(t *testing.T) {
	cfg := new(domain.Config)
	api := new(mockAPIClient)
	api.On("DeleteEnv", cfg, "NAME").Return(nil)

	r := &EnvsRepository{apiClient: api}

	err := r.Delete(cfg, "NAME")
	assert.Nil(t, err)
	api.AssertNumberOfCalls(t, "DeleteEnv", 1)
}

func TestEnvsRepository_Delete_ReturnErrorWhenDeleteEnvFailed(t *testing.T) {
	cfg := new(domain.Config)
	api := new(mockAPIClient)
	api.On("DeleteEnv", cfg, "NAME").Return(errors.New("SOMETHING_WRONG"))

	r := &EnvsRepository{apiClient: api}

	err := r.Delete(cfg, "NAME")
	assert.EqualError(t, err, "SOMETHING_WRONG")
	api.AssertNumberOfCalls(t, "DeleteEnv", 1)
}
