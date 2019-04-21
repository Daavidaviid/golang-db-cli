package cli

import (
	"github.com/urfave/cli"
)

var App *cli.App

func Init() error {
	App = cli.NewApp()
	return nil
}
