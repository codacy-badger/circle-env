package domain

type VCS string

const (
	GitHub    VCS = "github"
	BitBucket     = "bitbucket"
)

func (v VCS) IsValid() bool {
	switch v {
	case GitHub, BitBucket:
		return true
	default:
		return false
	}
}
