package utils

import (
	"errors"
)

// Command ...
type Command string

const (
	// Init ...
	Init Command = "init"
	// Show ...
	Show = "show"
	// Sync ...
	Sync = "sync"
)

func commandFromString(s string) (Command, error) {
	switch s {
	case "init", "show", "sync":
		return (Command)(s), nil
	default:
		return "", errors.New("invalid command")
	}
}

// Options ...
type Options struct {
	Help      bool
	Version   bool
	Command   Command
	JSON      bool
	Delete    bool
	NoConfirm bool
}

// NewOptions ...
func NewOptions(args []string) *Options {
	opts := new(Options)

	if len(args) == 0 {
		opts.Help = true
		return opts
	}

	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			if opts.Version == false {
				opts.Help = true
			}
		case "-v", "--version":
			if opts.Help == false {
				opts.Version = true
			}
		case "--json":
			opts.JSON = true
		case "--delete":
			opts.Delete = true
		case "--no-confirm":
			opts.NoConfirm = true
		}
	}

	if opts.Version {
		return opts
	}

	cmd, err := commandFromString(args[0])
	if err != nil {
		opts.Help = true
		return opts
	}
	opts.Command = cmd

	return opts
}
