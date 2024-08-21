package main

//'package main' is always for a standalone executable probram
//(as opposed to a library)

import (
	"database/sql"
	"log"
	"net/http"

	"csweeklyproj/db"
	"csweeklyproj/queries"

	"github.com/gin-gonic/gin"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// takes in the request state (gin.context) aswell as the db
// connection (sql.DB) to then return index.html along with
// a query of the problems
func getIndexPage(c *gin.Context, database *sql.DB) {
	users, err := queries.QueryUsers(database)
	if err != nil {
		log.Fatal(err)
	}
	//gin will then combine index.html with the users object and then
	//send the result back to the server
	c.HTML(http.StatusOK, "index.tmpl", users)
}

func getUsers(c *gin.Context, db *sql.DB) {
	users, err := queries.QueryUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func main() {

	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	//ensure the db closes on exit
	defer database.Close()

	//initialize Gin Router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//make a callback to getIndexPage instead of passing in the function
	router.GET("/", func(c *gin.Context) {
		getIndexPage(c, database)
	})
	router.GET("/users", func(c *gin.Context) {
		getUsers(c, database)
	})

	//attach router to http.Server and start it aswell
	router.Run("localhost:8080")
}
