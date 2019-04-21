package migration

import (
	"log"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Running 006_dataseeding.up.sql...")

		migrationContent, err := MigrationBox.FindString("006_dataseeding_up.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(migrationContent)
		return err
	}, func(db migrations.DB) error {
		log.Println("Running 006_dataseeding.down.sql...")

		migrationContent, err := MigrationBox.FindString("006_dataseeding_up.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(migrationContent)
		return err
	})
}
