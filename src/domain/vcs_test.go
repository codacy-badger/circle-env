package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VCSFromString_ReturnVCS(t *testing.T) {
	testCases := []struct {
		str      string
		expected VCS
	}{
		{"github", GitHub},
		{"bitbucket", BitBucket},
	}

	for _, c := range testCases {
		vcs, err := VCSFromString(c.str)
		assert.Equal(t, c.expected, vcs)
		assert.Nil(t, err)
	}
}

func Test_VCSFromString_ReturnErrorWhenInvalidVCSType(t *testing.T) {
	_, err := VCSFromString("INVALID_VCS_TYPE")
	assert.EqualError(t, err, "invalid vcs type")
}
