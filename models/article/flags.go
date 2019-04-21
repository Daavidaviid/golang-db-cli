package article

import "github.com/urfave/cli"

// Flags are all the flags used by the cli
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "id",
		Usage: "user's id",
	},
	cli.StringFlag{
		Name:  "name",
		Usage: "article's name",
	},
	cli.StringFlag{
		Name:  "content",
		Usage: "article's content",
	},
	cli.StringFlag{
		Name:  "userId",
		Usage: "article's userId",
	},
	cli.StringFlag{
		Name:  "userEmail",
		Usage: "article's userEmail",
	},
}
