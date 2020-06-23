package main

import (
	"flag"

	"github.com/kou-pg-0131/circle-env/src/infrastructures"
	"github.com/kou-pg-0131/circle-env/src/interfaces/controllers"
	"github.com/kou-pg-0131/circle-env/src/utils"
)

func main() {
	/*
	 * args
	 */

	j := flag.Bool("json", false, "json") // TODO
	flag.Parse()
	args := flag.Args()
	if flag.NArg() != 1 {
		panic("usage") // TODO
	}

	/*
	 * infrastructures
	 */

	api := infrastructures.NewCircleCIAPIClient()

	/*
	 * controllers
	 */

	cc := controllers.NewConfigController()
	ec := controllers.NewEnvsController(api)

	/*
	 * commands
	 */

	switch args[0] {
	case "init":
		cc.Initialize()
	case "show":
		if err := ec.Show(*j); err != nil {
			utils.Fatal(err)
		}
	default:
		panic("usage") // TODO
	}
}
