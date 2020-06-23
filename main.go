package main

import (
	"flag"

	"github.com/kou-pg-0131/circle-env/src/infrastructures"
	"github.com/kou-pg-0131/circle-env/src/interfaces/controllers"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		panic("usage") // TODO
	}

	api := infrastructures.NewCircleCIAPIClient()

	cc := controllers.NewConfigController()
	ec := controllers.NewEnvsController(api)

	switch args[0] {
	case "init":
		cc.Initialize()
	case "show":
		ec.Show()
	default:
		panic("usage") // TODO
	}
}
