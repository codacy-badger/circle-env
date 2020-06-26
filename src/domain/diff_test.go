package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff_sort(t *testing.T) {
	ds := &Diffs{
		{Name: "BBB"},
		{Name: "CCC"},
		{Name: "AAA"},
	}
	ds.sort()

	assert.Equal(t, &Diffs{
		{Name: "AAA"},
		{Name: "BBB"},
		{Name: "CCC"},
	}, ds)
}
