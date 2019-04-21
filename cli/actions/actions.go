package actions

import (
	"blogCLI/database"
	"blogCLI/migration"
	"blogCLI/models/user"
	"errors"
	"os"

	"github.com/urfave/cli"
)

func beforeCommand(c *cli.Context) error {
	// CHECK IF CONFIG FILE EXIST
	if _, err := os.Stat(".skyblogrc"); os.IsNotExist(err) {
		return errors.New("Can't find .skyblogrc file !")
	}

	// INIT DATABASE
	err := database.Init()
	if err != nil {
		return err
	}

	// INIT MIGRATION
	err = migration.Migrate()

	return err
}

func afterCommand(c *cli.Context) error {
	// database.DB.Close()
	return nil
}

var Commands = []cli.Command{
	{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "init CLI with some variables",
		Action:  Init,
	},
	{
		Name:        "create",
		Aliases:     []string{"c"},
		Usage:       "create an entity",
		Subcommands: CreateCommands,
		Flags:       user.Flags,
		Before:      beforeCommand,
		After:       afterCommand,
	},
	{
		Name:        "read",
		Aliases:     []string{"r"},
		Usage:       "read an entity",
		Subcommands: ReadCommands,
		Before:      beforeCommand,
		After:       afterCommand,
	},
	{
		Name:        "update",
		Aliases:     []string{"u"},
		Usage:       "update an entity",
		Subcommands: UpdateCommands,
		Before:      beforeCommand,
		After:       afterCommand,
	},
	{
		Name:        "delete",
		Aliases:     []string{"d"},
		Usage:       "delete an entity",
		Subcommands: DeleteCommands,
		Before:      beforeCommand,
		After:       afterCommand,
	},
	{
		Name:        "list",
		Aliases:     []string{"l"},
		Usage:       "list all elements of an entity type",
		Subcommands: ListCommands,
		Before:      beforeCommand,
		After:       afterCommand,
	},
}
