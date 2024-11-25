package config

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	migration, err := migrate.New(
		"file://../../migration",
		"postgres://dshestapalau:12345@localhost:5432/gogotask-profile?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	step := 1

	for {
		err := migration.Steps(step)

		if err == migrate.ErrNoChange {
			log.Println("All migrations applied successfully")
			break
		}

		if err != nil {
			log.Printf("Migration failed at step %d: %v", step, err)
			rollbackErr := migration.Steps(-1)
			if rollbackErr != nil {
				log.Fatalf("Rollback for step %d failed: %v", step, rollbackErr)
			}
			log.Printf("Rollback for step %d completed", step)
			break
		}
		step++
	}

	defer func() {
		srcErr, dbErr := migration.Close()
		if srcErr != nil || dbErr != nil {
			log.Printf("Error closing migration resources: srcErr=%v, dbErr=%v", srcErr, dbErr)
		}
	}()
}
