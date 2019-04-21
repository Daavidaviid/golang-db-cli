package database

import (
	"blogCLI/models"

	"github.com/go-pg/pg"
)

var DB *pg.DB

// Init init connection to postgresQL db
func Init() error {
	var config models.PostgresQLConfig

	err := config.Load(".skyblogrc")
	if err != nil {
		return err
	}

	options := &pg.Options{
		User:     config.User,
		Addr:     config.Address + ":" + config.Port,
		Database: config.Database,
		Password: config.Password,
	}

	DB = pg.Connect(options)

	return nil
}
