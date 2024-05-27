package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrateSql "github.com/rubenv/sql-migrate"
)

var DBConnection *sql.DB

//go:embed migrations/*.sql
var dataMigrate embed.FS

func DatabaseMigrate(paramDB *sql.DB) {
	migrations := &migrateSql.EmbedFileSystemMigrationSource{
		FileSystem: dataMigrate,
		Root:       "migrations",
	}

	n, err := migrateSql.Exec(paramDB, "postgres", migrations, migrateSql.Up)
	if err != nil {
		panic(err)
	}
	fmt.Println("Applied", n, "migrations")

}
