package main

import (
	"flag"

	"github.com/kou-pg-0131/circle-env/src/interfaces/controllers"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		panic("usage") // TODO
	}

	cc := controllers.NewConfigController()

	switch args[0] {
	case "init":
		cc.Initialize()
	default:
		panic("usage") // TODO
	}
}
