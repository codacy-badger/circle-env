package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kou-pg-0131/circle-env/src/infrastructures"
	"github.com/kou-pg-0131/circle-env/src/interfaces/controllers"
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
	de := infrastructures.NewDotenv()

	/*
	 * controllers
	 */

	cc := controllers.NewConfigController()
	ec := controllers.NewEnvsController(api, de)

	/*
	 * commands
	 */

	switch args[0] {
	case "init":
		if err := cc.Initialize(); err != nil {
			fatal(err)
		}
	case "show":
		if err := ec.Show(*j); err != nil {
			fatal(err)
		}
	case "push":
		if err := ec.Push(); err != nil {
			fatal(err)
		}
	default:
		panic("usage") // TODO
	}
}

func fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
