package main

import (
	"github.com/stretchr/testify/mock"
)

type mockConfigController struct {
	mock.Mock
}

func (m *mockConfigController) Initialize() error {
	args := m.Called()
	return args.Error(0)
}

type mockEnvsController struct {
	mock.Mock
}

func (m *mockEnvsController) Show(j bool) error {
	args := m.Called(j)
	return args.Error(0)
}

func (m *mockEnvsController) Sync(del, noconf bool) error {
	args := m.Called(del, noconf)
	return args.Error(0)
}
