package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValid_ReturnTrueWhenValid(t *testing.T) {
	vs := []VCS{
		"github",
		"bitbucket",
	}

	for _, v := range vs {
		assert.Equal(t, true, v.IsValid())
	}
}

func Test_IsValid_ReturnFalseWhenInvalid(t *testing.T) {
	assert.Equal(t, false, (VCS)("INVALID_VCS").IsValid())
}
