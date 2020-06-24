package utils

import (
	"errors"
	"os"
)

type Command string

const (
	Init Command = "init"
	Show         = "show"
	Push         = "push"
)

func commandFromString(s string) (Command, error) {
	switch s {
	case "init", "show", "push":
		return (Command)(s), nil
	default:
		return "", errors.New("invalid command")
	}
}

type options struct {
	Command Command
	Help    bool
	JSON    bool
}

func NewOptions() *options {
	opts := new(options)

	cmd, err := commandFromString(os.Args[1])
	if err != nil {
		Usage()
	}
	opts.Command = cmd

	for _, arg := range os.Args {
		switch arg {
		case "-h", "--help":
			Usage()
		case "--json":
			if opts.Command != Show {
				Usage()
			}
			opts.JSON = true
		}
	}

	return opts
}
