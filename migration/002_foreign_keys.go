package migration

import (
	"log"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Running 002_foreign_keys.up.sql...")

		migrationContent, err := MigrationBox.FindString("002_foreign_keys_up.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(migrationContent)
		return err
	}, func(db migrations.DB) error {
		log.Println("Running 002_foreign_keys.down.sql...")

		migrationContent, err := MigrationBox.FindString("002_foreign_keys_down.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(migrationContent)
		return err
	})
}
