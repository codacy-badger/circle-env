package utils

import (
	"fmt"
)

const (
	initDescription = "Interactively create the required configuration files."
	showDescription = "Output a list of environment variables set in the circleci project."
	syncDescription = "Syncs the value of `.circle-env/.env` with the environment variables of the circleci project."
)

type usage struct{}

func PrintUsage(cmd ...Command) {
	if cmd == nil {
		printBasicUsage()
	} else {
		switch cmd[0] {
		case Init:
			printInitUsage()
		case Show:
			printShowUsage()
		case Sync:
			printSyncUsage()
		default:
			printBasicUsage()
		}
	}
}

func printBasicUsage() {
	fmt.Printf(`Usage: circle-env <command> [options]

Commands:
  init     %s
  show     %s
  sync     %s
`, initDescription, showDescription, syncDescription)
}

func printInitUsage() {
	fmt.Printf(`Usage: circle-env init

  %s
`, initDescription)
}

func printShowUsage() {
	fmt.Printf(`Usage: circle-env show [options]

  %s

Options:
  --json     Output in json format.
`, showDescription)
}

func printSyncUsage() {
	fmt.Printf(`Usage: circle-env sync

  %s
`, syncDescription)
}
