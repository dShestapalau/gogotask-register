package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection(config Configuration) *gorm.DB {
	dsn := createDsn(config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func createDsn(config Configuration) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
}
