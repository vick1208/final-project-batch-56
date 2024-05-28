package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var DBConnection *sql.DB

//go:embed migrations/*.sql
var dataMigrate embed.FS

func DatabaseMigrate(DBParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dataMigrate,
		Root:       "migrations",
	}

	n, err := migrate.Exec(DBParam, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Println("Applied", n, "migrations")

}
