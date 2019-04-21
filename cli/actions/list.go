package actions

import (
	"blogCLI/models/article"
	"blogCLI/models/user"

	"github.com/urfave/cli"
)

var ListCommands = []cli.Command{
	{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "list an user",
		Action:  user.List,
		Flags:   user.Flags,
	},
	{
		Name:    "article",
		Aliases: []string{"a"},
		Usage:   "list an article",
		Action:  article.List,
		Flags:   article.Flags,
	},
}
