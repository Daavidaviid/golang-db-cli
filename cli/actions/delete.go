package actions

import (
	"blogCLI/models/article"
	"blogCLI/models/user"

	"github.com/urfave/cli"
)

var DeleteCommands = []cli.Command{
	{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "delete an user",
		Action:  user.Delete,
		Flags:   user.Flags,
	},
	{
		Name:    "article",
		Aliases: []string{"a"},
		Usage:   "delete an article",
		Action:  article.Delete,
		Flags:   article.Flags,
	},
}
