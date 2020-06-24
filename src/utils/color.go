package utils

import "fmt"

type color struct {
	str string
}

func Color(s string) *color {
	return &color{str: s}
}

func Colorf(format string, v ...interface{}) *color {
	return &color{str: fmt.Sprintf(format, v...)}
}

func (c *color) String() string {
	return fmt.Sprintf("%s\x1b[0m", c.str)
}

func (c *color) Bold() *color {
	return &color{str: fmt.Sprintf("\x1b[1m%s", c.str)}
}

func (c *color) Green() *color {
	return &color{str: fmt.Sprintf("\x1b[32m%s", c.str)}
}
