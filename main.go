package main

import (
	"database/sql"
	"fmt"
	"os"
	"project-indekost/database"
	"project-indekost/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed to load file env")
	}
	fmt.Println("success to load file env")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	err = DB.Ping()

	if err != nil {
		fmt.Println("Failed connected to database")
		panic(err)
	}
	fmt.Println("Successfully connected to database")

	database.DatabaseMigrate(DB)

	PORT := ":8080"
	server := routes.MainServer()
	server.Run(PORT)
}
