package main

import (
	"fmt"
	"os"

	"github.com/kou-pg-0131/circle-env/src/infrastructures"
	"github.com/kou-pg-0131/circle-env/src/interfaces/controllers"
	"github.com/kou-pg-0131/circle-env/src/utils"
)

func main() {
	app := NewApp()
	code, err := app.Run(os.Args[1:])
	if err != nil {
		app.PrintError(err)
	}

	os.Exit(code)
}

// IApp ...
type IApp interface {
	Run(args []string) (int, error)
	PrintError(err error)
}

// App ...
type App struct {
	configController controllers.IConfigController
	envsController   controllers.IEnvsController
}

// NewApp ...
func NewApp() *App {
	fs := infrastructures.NewFileSystem()
	api := infrastructures.NewCircleCIAPIClient()
	de := infrastructures.NewDotenv()

	return &App{
		configController: controllers.NewConfigController(fs),
		envsController:   controllers.NewEnvsController(api, fs, de),
	}
}

// Run ...
func (a *App) Run(args []string) (int, error) {
	opts := utils.NewOptions(args)

	if opts.Help {
		utils.PrintUsage(opts.Command)
		return 1, nil
	}

	if opts.Version {
		fmt.Println(utils.VERSION)
		return 0, nil
	}

	switch opts.Command {
	case utils.Init:
		if err := a.configController.Initialize(); err != nil {
			a.PrintError(err)
			return 1, err
		}
	case utils.Show:
		if err := a.envsController.Show(opts.JSON); err != nil {
			a.PrintError(err)
			return 1, err
		}
	case utils.Sync:
		if err := a.envsController.Sync(opts.Delete, opts.NoConfirm); err != nil {
			a.PrintError(err)
			return 1, err
		}
	default:
		utils.PrintUsage()
		return 1, nil
	}

	return 0, nil
}

// PrintError ...
func (a *App) PrintError(err error) {
	fmt.Println("Error:", err)
}
