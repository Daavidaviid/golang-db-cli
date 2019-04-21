package actions

import (
	"blogCLI/models/article"
	"blogCLI/models/user"

	"github.com/urfave/cli"
)

var ReadCommands = []cli.Command{
	{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "read an user",
		Action:  user.Read,
		Flags:   user.Flags,
	},
	{
		Name:    "article",
		Aliases: []string{"a"},
		Usage:   "read an article",
		Action:  article.Read,
		Flags:   article.Flags,
	},
}
