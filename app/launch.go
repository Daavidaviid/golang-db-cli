package app

import (
	"blogCLI/cli"
	"blogCLI/cli/actions"
	"os"
)

func LaunchCLI() error {

	err := cli.Init()
	if err != nil {
		return err
	}

	cli.App.Name = "SKYBLOG"
	cli.App.Usage = "CRUD app to manage articles and users"
	cli.App.Copyright = "NASA (c)"
	cli.App.Version = "v0.0.0.1.12.1"
	cli.App.Commands = actions.Commands

	return cli.App.Run(os.Args)
}
