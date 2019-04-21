package actions

import (
	"blogCLI/models/article"
	"blogCLI/models/user"

	"github.com/urfave/cli"
)

var UpdateCommands = []cli.Command{
	{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "update an user",
		Action:  user.Update,
		Flags:   user.Flags,
	},
	{
		Name:    "article",
		Aliases: []string{"a"},
		Usage:   "update an article",
		Action:  article.Update,
		Flags:   article.Flags,
	},
}
