package actions

import (
	"blogCLI/models/article"
	"blogCLI/models/user"

	"github.com/urfave/cli"
)

var CreateCommands = []cli.Command{
	{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "create an user",
		Action:  user.Create,
		Flags:   user.Flags,
	},
	{
		Name:    "article",
		Aliases: []string{"a"},
		Usage:   "create an article",
		Action:  article.Create,
		Flags:   article.Flags,
	},
}
