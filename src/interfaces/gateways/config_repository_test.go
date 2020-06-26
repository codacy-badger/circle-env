package gateways

import (
	"errors"
	"testing"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewConfigRepository()
 */

func Test_NewConfigRepository(t *testing.T) {
	fs := new(mockFileSystem)

	r := NewConfigRepository(fs)
	assert.Equal(t, &ConfigRepository{fs: fs}, r)
}

/*
 * ConfigRepository.Save()
 */

func TestConfigRepository_Save_ReturnNil(t *testing.T) {
	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}

	cf := new(mockFile)
	cf.On("Write", []byte("vcs = github\nuser = USER\nrepo = REPO")).Return(0, nil)
	cf.On("Close").Return(nil)

	tf := new(mockFile)
	tf.On("Write", []byte("TOKEN")).Return(0, nil)
	tf.On("Close").Return(nil)

	fs := new(mockFileSystem)
	fs.On("Mkdir", ".circle-env").Return(nil)
	fs.On("Create", ".circle-env/config").Return(cf, nil)
	fs.On("Create", ".circle-env/circle-token").Return(tf, nil)

	r := &ConfigRepository{fs: fs}

	err := r.Save(cfg)
	assert.Nil(t, err)
	cf.AssertNumberOfCalls(t, "Write", 1)
	cf.AssertNumberOfCalls(t, "Close", 1)
	tf.AssertNumberOfCalls(t, "Write", 1)
	tf.AssertNumberOfCalls(t, "Close", 1)
	fs.AssertNumberOfCalls(t, "Mkdir", 1)
	fs.AssertNumberOfCalls(t, "Create", 2)
}

func TestConfigRepository_Save_ReturnErrorWhenCreateDirFailed(t *testing.T) {
	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}

	fs := new(mockFileSystem)
	fs.On("Mkdir", ".circle-env").Return(errors.New("SOMETHING_WRONG"))

	r := &ConfigRepository{fs: fs}

	err := r.Save(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	fs.AssertNumberOfCalls(t, "Mkdir", 1)
}

func TestConfigRepository_Save_ReturnErrorWhenCreateConfigFileFailed(t *testing.T) {
	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}

	fs := new(mockFileSystem)
	fs.On("Mkdir", ".circle-env").Return(nil)
	fs.On("Create", ".circle-env/config").Return((*mockFile)(nil), errors.New("SOMETHING_WRONG"))

	r := &ConfigRepository{fs: fs}

	err := r.Save(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	fs.AssertNumberOfCalls(t, "Mkdir", 1)
	fs.AssertNumberOfCalls(t, "Create", 1)
}

func TestConfigRepository_Save_ReturnErrorWhenWriteConfigFileFailed(t *testing.T) {
	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}

	cf := new(mockFile)
	cf.On("Write", []byte("vcs = github\nuser = USER\nrepo = REPO")).Return(0, errors.New("SOMETHING_WRONG"))
	cf.On("Close").Return(nil)

	fs := new(mockFileSystem)
	fs.On("Mkdir", ".circle-env").Return(nil)
	fs.On("Create", ".circle-env/config").Return(cf, nil)

	r := &ConfigRepository{fs: fs}

	err := r.Save(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	cf.AssertNumberOfCalls(t, "Write", 1)
	cf.AssertNumberOfCalls(t, "Close", 1)
	fs.AssertNumberOfCalls(t, "Mkdir", 1)
	fs.AssertNumberOfCalls(t, "Create", 1)
}

func TestConfigRepository_Save_ReturnErrorWhenCreateTokenFileFailed(t *testing.T) {
	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}

	cf := new(mockFile)
	cf.On("Write", []byte("vcs = github\nuser = USER\nrepo = REPO")).Return(0, nil)
	cf.On("Close").Return(nil)

	fs := new(mockFileSystem)
	fs.On("Mkdir", ".circle-env").Return(nil)
	fs.On("Create", ".circle-env/config").Return(cf, nil)
	fs.On("Create", ".circle-env/circle-token").Return((*mockFile)(nil), errors.New("SOMETHING_WRONG"))

	r := &ConfigRepository{fs: fs}

	err := r.Save(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	cf.AssertNumberOfCalls(t, "Write", 1)
	cf.AssertNumberOfCalls(t, "Close", 1)
	fs.AssertNumberOfCalls(t, "Mkdir", 1)
	fs.AssertNumberOfCalls(t, "Create", 2)
}

func TestConfigRepository_Save_ReturnErrorWhenWriteTokenFileFailed(t *testing.T) {
	cfg := &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}

	cf := new(mockFile)
	cf.On("Write", []byte("vcs = github\nuser = USER\nrepo = REPO")).Return(0, nil)
	cf.On("Close").Return(nil)

	tf := new(mockFile)
	tf.On("Write", []byte("TOKEN")).Return(0, errors.New("SOMETHING_WRONG"))
	tf.On("Close").Return(nil)

	fs := new(mockFileSystem)
	fs.On("Mkdir", ".circle-env").Return(nil)
	fs.On("Create", ".circle-env/config").Return(cf, nil)
	fs.On("Create", ".circle-env/circle-token").Return(tf, nil)

	r := &ConfigRepository{fs: fs}

	err := r.Save(cfg)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	cf.AssertNumberOfCalls(t, "Write", 1)
	cf.AssertNumberOfCalls(t, "Close", 1)
	tf.AssertNumberOfCalls(t, "Write", 1)
	tf.AssertNumberOfCalls(t, "Close", 1)
	fs.AssertNumberOfCalls(t, "Mkdir", 1)
	fs.AssertNumberOfCalls(t, "Create", 2)
}

/*
 * ConfigRepository.Get()
 */

func TestConfigRepository_Get_ReturnConfig(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(true)
	fs.On("IsExists", ".circle-env/circle-token").Return(true)
	fs.On("Read", ".circle-env/config").Return([]byte("vcs = github\nuser = USER\nrepo = REPO"), nil)
	fs.On("Read", ".circle-env/circle-token").Return([]byte("TOKEN"), nil)

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.Equal(t, &domain.Config{
		VCS:   domain.GitHub,
		User:  "USER",
		Repo:  "REPO",
		Token: "TOKEN",
	}, cfg)
	assert.Nil(t, err)
	fs.AssertNumberOfCalls(t, "IsExists", 2)
	fs.AssertNumberOfCalls(t, "Read", 2)
}

func TestConfigRepository_Get_ReturnErrorWhenConfigFileNotFound(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(false)

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.EqualError(t, err, "`.circle-env/config` not found, run `circle-env init`")
	assert.Nil(t, cfg)
	fs.AssertNumberOfCalls(t, "IsExists", 1)
}

func TestConfigRepository_Get_ReturnErrorWhenTokenFileNotFound(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(true)
	fs.On("IsExists", ".circle-env/circle-token").Return(false)

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.EqualError(t, err, "`.circle-env/circle-token` not found, run `circle-env init`")
	assert.Nil(t, cfg)
	fs.AssertNumberOfCalls(t, "IsExists", 2)
}

func TestConfigRepository_Get_ReturnErrorWhenReadConfigFileFailed(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(true)
	fs.On("IsExists", ".circle-env/circle-token").Return(true)
	fs.On("Read", ".circle-env/config").Return(([]byte)(nil), errors.New("SOMETHING_WRONG"))

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.EqualError(t, err, "SOMETHING_WRONG")
	assert.Nil(t, cfg)
	fs.AssertNumberOfCalls(t, "IsExists", 2)
	fs.AssertNumberOfCalls(t, "Read", 1)
}

func TestConfigRepository_Get_ReturnErrorWhenLoadConfigFileFailed(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(true)
	fs.On("IsExists", ".circle-env/circle-token").Return(true)
	fs.On("Read", ".circle-env/config").Return([]byte("invalid_ini"), nil)

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.NotNil(t, err)
	assert.Nil(t, cfg)
	fs.AssertNumberOfCalls(t, "IsExists", 2)
	fs.AssertNumberOfCalls(t, "Read", 1)
}

func TestConfigRepository_Get_ReturnErrorWhenInvalidVCSType(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(true)
	fs.On("IsExists", ".circle-env/circle-token").Return(true)
	fs.On("Read", ".circle-env/config").Return([]byte("vcs = invalid_vcs\nuser = USER\nrepo = REPO"), nil)

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.EqualError(t, err, "`invalid_vcs` is invalid vcs type, please check `.circle-env/config`")
	assert.Nil(t, cfg)
	fs.AssertNumberOfCalls(t, "IsExists", 2)
	fs.AssertNumberOfCalls(t, "Read", 1)
}

func TestConfigRepository_Get_ReturnErrorWhenReadTokenFileFailed(t *testing.T) {
	fs := new(mockFileSystem)
	fs.On("IsExists", ".circle-env/config").Return(true)
	fs.On("IsExists", ".circle-env/circle-token").Return(true)
	fs.On("Read", ".circle-env/config").Return([]byte("vcs = github\nuser = USER\nrepo = REPO"), nil)
	fs.On("Read", ".circle-env/circle-token").Return(([]byte)(nil), errors.New("SOMETHING_WRONG"))

	r := &ConfigRepository{fs: fs}

	cfg, err := r.Get()
	assert.EqualError(t, err, "SOMETHING_WRONG")
	assert.Nil(t, cfg)
	fs.AssertNumberOfCalls(t, "IsExists", 2)
	fs.AssertNumberOfCalls(t, "Read", 2)
}
