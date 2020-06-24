package utils

import (
	"fmt"
	"os"
)

const (
	initDescription = "Interactively create the required configuration files."
	showDescription = "Output a list of environment variables set in the circleci project."
	syncDescription = "Syncs the value of `.circle-env/.env` with the environment variables of the circleci project."
)

type usage struct{}

var Usage = new(usage)

func (u *usage) Print(cmd ...Command) {
	if cmd == nil {
		u.printBasicUsage()
	} else {
		switch cmd[0] {
		case Init:
			u.printInitUsage()
		case Show:
			u.printShowUsage()
		case Sync:
			u.printSyncUsage()
		}
	}

	fmt.Println("")
	os.Exit(1)
}

func (u *usage) printBasicUsage() {
	fmt.Printf(`Usage: circle-env <command> [options]

Commands:
  init     %s
  show     %s
  sync     %s
`, initDescription, showDescription, syncDescription)
}

func (u *usage) printInitUsage() {
	fmt.Printf(`Usage: circle-env init

  %s
`, initDescription)
}

func (u *usage) printShowUsage() {
	fmt.Printf(`Usage: circle-env show [options]

  %s

Options:
  --json     Output in json format.
`, showDescription)
}

func (u *usage) printSyncUsage() {
	fmt.Printf(`Usage: circle-env sync

  %s
`, syncDescription)
}
