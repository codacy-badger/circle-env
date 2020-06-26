package domain

import "errors"

// VCS ...
type VCS string

const (
	// GitHub ...
	GitHub VCS = "github"
	// BitBucket ...
	BitBucket = "bitbucket"
)

// VCSFromString ...
func VCSFromString(s string) (VCS, error) {
	switch s {
	case "github", "bitbucket":
		return (VCS)(s), nil
	default:
		return "", errors.New("invalid vcs type")
	}
}
