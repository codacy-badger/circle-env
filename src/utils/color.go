package utils

import "fmt"

type Color struct {
	str string
}

func (c *Color) color(code int, s string) string {
	return fmt.Sprintf("\x1b[%dm%s", code, s)
}

func (c *Color) clear() *Color {
	return &Color{str: fmt.Sprintf("%s\x1b[0m", c.str)}
}

func Colorf(format string, v ...interface{}) *Color {
	return &Color{str: fmt.Sprintf(format, v...)}
}

func (c *Color) String() string {
	return c.clear().str
}

func (c *Color) Bold() *Color {
	return &Color{str: c.color(1, c.str)}
}

func (c *Color) Secondary() *Color {
	return &Color{str: c.color(2, c.str)}
}

func (c *Color) Green() *Color {
	return &Color{str: c.color(32, c.str)}
}

func (c *Color) Red() *Color {
	return &Color{str: c.color(31, c.str)}
}
