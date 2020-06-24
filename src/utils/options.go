package utils

import (
	"errors"
	"os"
)

type Command string

const (
	Init Command = "init"
	Show         = "show"
	Sync         = "sync"
)

func commandFromString(s string) (Command, error) {
	switch s {
	case "init", "show", "sync":
		return (Command)(s), nil
	default:
		return "", errors.New("invalid command")
	}
}

type options struct {
	Command Command
	JSON    bool
}

func NewOptions() *options {
	opts := new(options)

	if len(os.Args) < 2 {
		Usage.Print()
	}

	cmd, err := commandFromString(os.Args[1])
	if err != nil {
		Usage.Print()
	}
	opts.Command = cmd

	for _, arg := range os.Args {
		switch arg {
		case "-h", "--help":
			Usage.Print(opts.Command)
		case "--json":
			if opts.Command != Show {
				Usage.Print()
			}
			opts.JSON = true
		}
	}

	return opts
}
