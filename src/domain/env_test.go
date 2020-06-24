package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvs_Has_ReturnTrue(t *testing.T) {
	es := &Envs{
		{Name: "HOGE", Value: "PIYO"},
		{Name: "FOO", Value: "BAR"},
	}

	e := &Env{Name: "HOGE", Value: "BAR"}

	assert.Equal(t, true, es.Has(e))
}

func TestEnvs_Has_ReturnFalse(t *testing.T) {
	es := &Envs{
		{Name: "FOO", Value: "BAR"},
	}

	e := &Env{Name: "HOGE", Value: "BAR"}

	assert.Equal(t, false, es.Has(e))
}
