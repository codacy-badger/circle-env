package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvs_Has(t *testing.T) {
	es := &Envs{
		{Name: "HOGE", Value: "PIYO"},
		{Name: "FOO", Value: "BAR"},
	}

	assert.Equal(t, true, es.Has("HOGE"))
	assert.Equal(t, true, es.Has("FOO"))
	assert.Equal(t, false, es.Has("PIYO"))
	assert.Equal(t, false, es.Has("BAR"))
}

func TestEnvs_Sort(t *testing.T) {
	es := &Envs{
		{Name: "BBB"},
		{Name: "CCC"},
		{Name: "AAA"},
	}

	es.Sort()
	assert.Equal(t, &Envs{
		{Name: "AAA"},
		{Name: "BBB"},
		{Name: "CCC"},
	}, es)
}
