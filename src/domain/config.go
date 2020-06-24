package domain

import (
	"fmt"
	"strings"
)

type Config struct {
	VCS   VCS
	User  string
	Repo  string
	Token string
}

func (c *Config) Ini() string {
	return strings.Join([]string{
		fmt.Sprintf("vcs = %s", c.VCS),
		fmt.Sprintf("user = %s", c.User),
		fmt.Sprintf("repo = %s", c.Repo),
	}, "\n")
}

func (c *Config) Slug() string {
	return fmt.Sprintf("%s/%s/%s", c.VCS, c.User, c.Repo)
}
