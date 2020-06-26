package infrastructures

import (
	"github.com/stretchr/testify/mock"
)

type mockHTTPClient struct {
	mock.Mock
}

func (m *mockHTTPClient) Get(url string, header map[string]string) (*HTTPResponse, error) {
	args := m.Called(url, header)
	return args.Get(0).(*HTTPResponse), args.Error(1)
}

func (m *mockHTTPClient) Post(url string, header map[string]string, body []byte) (*HTTPResponse, error) {
	args := m.Called(url, header, body)
	return args.Get(0).(*HTTPResponse), args.Error(1)
}

func (m *mockHTTPClient) Delete(url string, header map[string]string) (*HTTPResponse, error) {
	args := m.Called(url, header)
	return args.Get(0).(*HTTPResponse), args.Error(1)
}
