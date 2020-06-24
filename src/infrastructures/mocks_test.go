package infrastructures

import (
	"github.com/stretchr/testify/mock"
)

type mockHttpClient struct {
	mock.Mock
}

func (m *mockHttpClient) Get(url string, header map[string]string) (*HttpResponse, error) {
	args := m.Called(url, header)
	return args.Get(0).(*HttpResponse), args.Error(1)
}

func (m *mockHttpClient) Post(url string, header map[string]string, body []byte) (*HttpResponse, error) {
	args := m.Called(url, header, body)
	return args.Get(0).(*HttpResponse), args.Error(1)
}
