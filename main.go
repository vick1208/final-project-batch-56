package main

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "admin"
	dbname   = "db_boarding"
)

var (
	Database *sql.DB
	err      error
)

func main() {

	err = godotenv.Load("config/.env.local")
	if err != nil {
		fmt.Println("failed to load file env: ", err)
	}
	fmt.Println("success to load file env")

	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
	fmt.Println("Successfully connected to database")
}
