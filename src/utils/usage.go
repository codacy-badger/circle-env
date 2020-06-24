package utils

import (
	"fmt"
	"os"
)

const (
	initDescription = "Interactively create the required configuration files."
	showDescription = "Output a list of environment variables set in the circleci project."
	pushDescription = "Syncs the value of `.circle-env/.env` with the environment variables of the circleci project."
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
		case Push:
			u.printPushUsage()
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
  push     %s
`, initDescription, showDescription, pushDescription)
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

func (u *usage) printPushUsage() {
	fmt.Printf(`Usage: circle-env push

  %s
`, pushDescription)
}
