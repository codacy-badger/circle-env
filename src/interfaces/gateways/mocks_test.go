package gateways

import (
	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/stretchr/testify/mock"
)

type mockAPIClient struct {
	mock.Mock
}

func (m *mockAPIClient) GetEnvs(cfg *domain.Config) (*domain.Envs, error) {
	args := m.Called(cfg)
	return args.Get(0).(*domain.Envs), args.Error(1)
}

func (m *mockAPIClient) CreateEnv(cfg *domain.Config, e *domain.Env) error {
	args := m.Called(cfg, e)
	return args.Error(0)
}

func (m *mockAPIClient) DeleteEnv(cfg *domain.Config, name string) error {
	args := m.Called(cfg, name)
	return args.Error(0)
}

type mockDotenv struct {
	mock.Mock
}

func (m *mockDotenv) Load(path string) (*domain.Envs, error) {
	args := m.Called(path)
	return args.Get(0).(*domain.Envs), args.Error(1)
}

type mockFileSystem struct {
	mock.Mock
}

func (m *mockFileSystem) Read(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockFileSystem) IsExists(path string) bool {
	args := m.Called(path)
	return args.Bool(0)
}

func (m *mockFileSystem) Mkdir(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *mockFileSystem) Create(path string) (IFile, error) {
	args := m.Called(path)
	return args.Get(0).(IFile), args.Error(1)
}

type mockFile struct {
	mock.Mock
}

func (m *mockFile) Write(b []byte) (int, error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}

func (m *mockFile) Close() error {
	args := m.Called()
	return args.Error(0)
}
