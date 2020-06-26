package utils

import (
	"errors"
	"os"
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
func NewOptions() *Options {
	args := os.Args
	opts := new(Options)

	if len(args) < 2 {
		PrintUsage()
		os.Exit(1)
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
		PrintVersion()
		os.Exit(0)
	}

	cmd, err := commandFromString(args[1])
	if err != nil {
		PrintUsage()
		os.Exit(1)
	}

	opts.Command = cmd

	if opts.Help {
		PrintUsage(opts.Command)
		os.Exit(1)
	}

	return opts
}
