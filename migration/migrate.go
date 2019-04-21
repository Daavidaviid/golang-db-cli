package migration

import (
	"blogCLI/database"
	"log"

	"github.com/go-pg/migrations"
	packr "github.com/gobuffalo/packr/v2"
)

var MigrationBox = packr.New("Migrations Box", "../migrationsFiles")

// Migrate to run migrations defined in sql files
func Migrate() error {

	oldVersion, newVersion, err := migrations.Run(database.DB, "init")
	oldVersion, newVersion, err = migrations.Run(database.DB)
	if err != nil {
		return err
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	}
	// else {
	// 	log.Printf("version is %d\n", oldVersion)
	// }
	return nil
}
