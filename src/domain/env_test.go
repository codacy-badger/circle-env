package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Envs.Has()
 */

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

/*
 * Envs.Get()
 */

func TestEnvs_Get(t *testing.T) {
	es := &Envs{
		{Name: "HOGE", Value: "PIYO"},
		{Name: "FOO", Value: "BAR"},
	}

	assert.Equal(t, &Env{Name: "HOGE", Value: "PIYO"}, es.Get("HOGE"))
	assert.Equal(t, &Env{Name: "FOO", Value: "BAR"}, es.Get("FOO"))
	assert.Nil(t, es.Get("A"))
}

/*
 * Envs.Sort()
 */

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
