package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Ini_ReturnIniString(t *testing.T) {
	c := &Config{
		VCS:   GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}

	assert.Equal(t, `vcs = github
user = user
repo = repo`, c.Ini())
}

func TestConfig_Slug(t *testing.T) {
	c := &Config{
		VCS:   GitHub,
		User:  "user",
		Repo:  "repo",
		Token: "token",
	}

	assert.Equal(t, "github/user/repo", c.Slug())
}
