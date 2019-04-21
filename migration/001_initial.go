package migration

import (
	"log"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Running 001_initial.up.sql...")

		migrationContent, err := MigrationBox.FindString("001_initial_up.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(migrationContent)
		return err
	}, func(db migrations.DB) error {
		log.Println("Running 001_initial.down.sql...")

		migrationContent, err := MigrationBox.FindString("001_initial_down.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(migrationContent)
		return err
	})
}
