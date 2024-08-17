package main

//'package main' is always for a standalone executable probram
//(as opposed to a library)

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type problem struct {
	ID          string ` json:"id" `
	Text        string ` json:"text" `
	Hint        string ` json:"hint" `
	Constraints string ` json:"constraints" `
	Solution    string ` json:"solution" `
	IsProject   bool   ` json:"isproject" `
}

func getIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func returnENV() {

}

// does not take in anything, returns db and err
func initDB() (*sql.DB, error) {
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

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	//ensure the db closes on exit
	defer db.Close()

	//initialize Gin Router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndexPage)
	router.GET("/router")

	//attach router to http.Server and start it aswell
	router.Run("localhost:8080")
}
