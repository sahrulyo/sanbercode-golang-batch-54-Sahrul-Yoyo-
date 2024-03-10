package database

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "sql_migrations", // Sesuaikan dengan nama direktori migrasi Anda
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	DbConnection = dbParam

	fmt.Println("Applied", n, "migrations!")
}
