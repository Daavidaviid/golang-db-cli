package user

import "github.com/urfave/cli"

// Flags are all the flags used by the cli
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "id",
		Usage: "user's id",
	},
	cli.StringFlag{
		Name:  "firstName",
		Usage: "user's first name",
	},
	cli.StringFlag{
		Name:  "lastName",
		Usage: "user's last name",
	},
	cli.StringFlag{
		Name:  "password",
		Usage: "user's password",
	},
	cli.StringFlag{
		Name:  "phone",
		Usage: "entity's phone",
	},
	cli.StringFlag{
		Name:  "email",
		Usage: "entity's email",
	},
}
