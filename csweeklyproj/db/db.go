package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitDB() (*sql.DB, error) {
	//initialize env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("TURSO_DATABASE_URL")
	if url == "" { //check if ENV is working correctly
		log.Fatal("TURSO_DATABASE_URL environment variable is not set")
	}

	//initialize turso db
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db  %s: %s", url, err)
		os.Exit(1)
		//"1" means error, "0" means all good
	}
	return db, err
}
