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

	api := infrastructures.NewCircleCIAPIClient()
	de := infrastructures.NewDotenv()

	/*
	 * controllers
	 */

	cc := controllers.NewConfigController()
	ec := controllers.NewEnvsController(api, de)

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
	case utils.Push:
		if err := ec.Push(); err != nil {
			fatal(err)
		}
	default:
		utils.Usage.Print()
	}
}

func fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
