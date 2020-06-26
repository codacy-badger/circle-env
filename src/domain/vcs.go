package domain

import "errors"

type VCS string

const (
	GitHub    VCS = "github"
	BitBucket     = "bitbucket"
)

func VCSFromString(s string) (VCS, error) {
	switch s {
	case "github", "bitbucket":
		return (VCS)(s), nil
	default:
		return "", errors.New("invalid vcs type")
	}
}
