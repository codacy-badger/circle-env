package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Colorf(t *testing.T) {
	c := Colorf("%s", "STRING")
	assert.Equal(t, &Color{str: "STRING"}, c)
}

func TestColor_String(t *testing.T) {
	s := Colorf("STRING").String()
	assert.Equal(t, "STRING\x1b[0m", s)
}

func TestColor_Bold(t *testing.T) {
	s := Colorf("STRING").Bold().String()
	assert.Equal(t, "\x1b[1mSTRING\x1b[0m", s)
}

func TestColor_Secondary(t *testing.T) {
	s := Colorf("STRING").Secondary().String()
	assert.Equal(t, "\x1b[2mSTRING\x1b[0m", s)
}

func TestColor_Green(t *testing.T) {
	s := Colorf("STRING").Green().String()
	assert.Equal(t, "\x1b[32mSTRING\x1b[0m", s)
}

func TestColor_Red(t *testing.T) {
	s := Colorf("STRING").Red().String()
	assert.Equal(t, "\x1b[31mSTRING\x1b[0m", s)
}
