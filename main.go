package main

import (
	"fmt"
	"os"

	"github.com/kou-pg-0131/circle-env/src/infrastructures"
	"github.com/kou-pg-0131/circle-env/src/interfaces/controllers"
	"github.com/kou-pg-0131/circle-env/src/utils"
)

func main() {
	/*
	 * args
	 */

	opts := utils.NewOptions()

	/*
	 * infrastructures
	 */

	fs := infrastructures.NewFileSystem()
	api := infrastructures.NewCircleCIAPIClient()
	de := infrastructures.NewDotenv()

	/*
	 * controllers
	 */

	cc := controllers.NewConfigController(fs)
	ec := controllers.NewEnvsController(api, fs, de)

	/*
	 * commands
	 */

	switch opts.Command {
	case utils.Init:
		if err := cc.Initialize(); err != nil {
			fatal(err)
		}
	case utils.Show:
		if err := ec.Show(opts.JSON); err != nil {
			fatal(err)
		}
	case utils.Sync:
		if err := ec.Sync(opts.Delete, opts.NoConfirm); err != nil {
			fatal(err)
		}
	default:
		utils.PrintUsage()
		os.Exit(1)
	}
}

func fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
