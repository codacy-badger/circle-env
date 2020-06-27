package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * NewApp()
 */

func Test_NewApp(t *testing.T) {
	a := NewApp()
	assert.NotNil(t, a.configController)
	assert.NotNil(t, a.envsController)
}

/*
 * App.Run()
 */

func TestApp_Run_ReturnOneWhenInvalidArgs(t *testing.T) {
	a := new(App)

	argsList := [][]string{
		{},
		{"foo", "bar"},
	}

	for _, args := range argsList {
		code, err := a.Run(args)
		assert.Equal(t, 1, code)
		assert.Nil(t, err)
	}
}

func TestApp_Run_Help_ReturnOne(t *testing.T) {
	a := new(App)

	argsList := [][]string{
		{"-h"},
		{"--help"},
	}

	for _, args := range argsList {
		code, err := a.Run(args)
		assert.Equal(t, 1, code)
		assert.Nil(t, err)
	}
}

func TestApp_Run_Version_ReturnZero(t *testing.T) {
	a := new(App)

	argsList := [][]string{
		{"-v"},
		{"--version"},
	}

	for _, args := range argsList {
		code, err := a.Run(args)
		assert.Equal(t, 0, code)
		assert.Nil(t, err)
	}
}

func TestApp_Run_Init_ReturnZero(t *testing.T) {
	cc := new(mockConfigController)
	cc.On("Initialize").Return(nil)

	a := &App{configController: cc}

	code, err := a.Run([]string{"init"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	cc.AssertNumberOfCalls(t, "Initialize", 1)
}

func TestApp_Run_Init_ReturnErrorWhenInitializeFailed(t *testing.T) {
	cc := new(mockConfigController)
	cc.On("Initialize").Return(errors.New("SOMETHING_WRONG"))

	a := &App{configController: cc}

	code, err := a.Run([]string{"init"})
	assert.Equal(t, 1, code)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	cc.AssertNumberOfCalls(t, "Initialize", 1)
}

func TestApp_Run_Show_ReturnZero(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Show", false).Return(nil)

	a := &App{envsController: ec}

	code, err := a.Run([]string{"show"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	ec.AssertNumberOfCalls(t, "Show", 1)
}

func TestApp_Run_ShowWithJSON_ReturnZero(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Show", true).Return(nil)

	a := &App{envsController: ec}

	code, err := a.Run([]string{"show", "--json"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	ec.AssertNumberOfCalls(t, "Show", 1)
}

func TestApp_Run_Show_ReturnErrorWhenShowFailed(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Show", false).Return(errors.New("SOMETHING_WRONG"))

	a := &App{envsController: ec}

	code, err := a.Run([]string{"show"})
	assert.Equal(t, 1, code)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	ec.AssertNumberOfCalls(t, "Show", 1)
}

func TestApp_Run_Sync_ReturnZero(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Sync", false, false).Return(nil)

	a := &App{envsController: ec}

	code, err := a.Run([]string{"sync"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	ec.AssertNumberOfCalls(t, "Sync", 1)
}

func TestApp_Run_SyncWithDeleteAndNoConfirm_ReturnZero(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Sync", true, true).Return(nil)

	a := &App{envsController: ec}

	code, err := a.Run([]string{"sync", "--delete", "--no-confirm"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	ec.AssertNumberOfCalls(t, "Sync", 1)
}

func TestApp_Run_SyncWithDelete_ReturnZero(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Sync", true, false).Return(nil)

	a := &App{envsController: ec}

	code, err := a.Run([]string{"sync", "--delete"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	ec.AssertNumberOfCalls(t, "Sync", 1)
}

func TestApp_Run_SyncWithNoConfirm_ReturnZero(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Sync", false, true).Return(nil)

	a := &App{envsController: ec}

	code, err := a.Run([]string{"sync", "--no-confirm"})
	assert.Equal(t, 0, code)
	assert.Nil(t, err)
	ec.AssertNumberOfCalls(t, "Sync", 1)
}

func TestApp_Run_Sync_ReturnErrorSyncFailed(t *testing.T) {
	ec := new(mockEnvsController)
	ec.On("Sync", false, false).Return(errors.New("SOMETHING_WRONG"))

	a := &App{envsController: ec}

	code, err := a.Run([]string{"sync"})
	assert.Equal(t, 1, code)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	ec.AssertNumberOfCalls(t, "Sync", 1)
}
