package presenters

import (
	"testing"

	"github.com/kou-pg-0131/circle-env/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewEnvsPresenter()
 */

func Test_NewEnvsPresenter(t *testing.T) {
	jp := NewEnvsPresenter(true)
	pp := NewEnvsPresenter(false)

	assert.Equal(t, new(JSONEnvsPresenter), jp)
	assert.Equal(t, new(PlainEnvsPresenter), pp)
}

/*
 * PlainEnvsPresenter.String()
 */

func TestPlainEnvsPresenter_String(t *testing.T) {
	testCases := []struct {
		es       *domain.Envs
		expected string
	}{
		{&domain.Envs{}, "no environment variables are set."},
		{&domain.Envs{
			{Name: "HOGE", Value: "PIYO"},
			{Name: "FOO", Value: "BAR"},
		}, "HOGE = \"PIYO\"\nFOO = \"BAR\""},
	}

	for _, c := range testCases {
		s, err := new(PlainEnvsPresenter).String(c.es)
		assert.Equal(t, c.expected, s)
		assert.Nil(t, err)
	}
}

/*
 * JSONEnvsPresenter
 */

func TestJSONEnvsPresenter_String(t *testing.T) {
	es := &domain.Envs{
		{Name: "HOGE", Value: "PIYO"},
		{Name: "FOO", Value: "BAR"},
	}

	s, err := new(JSONEnvsPresenter).String(es)
	assert.Equal(t, `[
  {
    "name": "HOGE",
    "value": "PIYO"
  },
  {
    "name": "FOO",
    "value": "BAR"
  }
]`, s)
	assert.Nil(t, err)
}
