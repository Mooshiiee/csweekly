// This is the main entrypoint to the application.
// its only job is to process HTTP requests.
// IT SHOULD NOT contain application logic,
// that job is for Services package and Component Package
package main

//'package main' is always for a standalone executable program
//(as opposed to a library)

import (
	"database/sql"
	"log"
	"net/http"

	"csweeklyproj/db"
	"csweeklyproj/queries"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
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
	c.HTML(http.StatusOK, "index.html", users)
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

func getProblems(c *gin.Context, database *sql.DB) {
	problems, err := queries.QueryProblems(database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, problems)
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
	//initialize the templ html renderer
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTempleRenderer{
		fallbackHtmlRenderer: ginHtmlRenderer,
	}

	//Disable trusted proxy warning.
	router.SetTrustedProxies(nil)

	//make a callback to getIndexPage instead of
	//passing in the function
	router.GET("/", func(c *gin.Context) {
		getIndexPage(c, database)
	})
	router.GET("/users", func(c *gin.Context) {
		getUsers(c, database)
	})
	router.GET("/problems", func(c *gin.Context) {
		getProblems(c, database)
	})

	//attach router to http.Server and start it aswell
	router.Run("localhost:8080")
}
