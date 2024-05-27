package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "db_boarding_house"
)

var (
	Database *sql.DB
	err      error
)

func main() {
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Database, err = sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}
	defer Database.Close()
	err = Database.Ping()
	if err != nil {
		fmt.Println("Failed connected to database")
		panic(err)
	}
}
