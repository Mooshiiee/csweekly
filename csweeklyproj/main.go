package main

//'package main' is always for a standalone executable probram
//(as opposed to a library)

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type problem struct {
	ID          string ` json:"id" `
	Text        string ` json:"text" `
	Hint        string ` json:"hint" `
	Constraints string ` json:"constraints" `
	Solution    string ` json:"solution" `
	IsProject   bool   ` json:"isproject" `
}

var problems = []problem{
	{
		ID:          "1",
		Text:        "Implement a function to reverse a string.",
		Hint:        "Think about string indexing.",
		Constraints: "Input string length <= 1000.",
		Solution:    "Use slicing to reverse the string.",
		IsProject:   false,
	},
	{
		ID:          "2",
		Text:        "Find the maximum value in an array of integers.",
		Hint:        "Iterate through the array and keep track of the maximum value found.",
		Constraints: "Array length <= 10000.",
		Solution:    "Use a loop to compare each element with the current maximum.",
		IsProject:   false,
	},
	{
		ID:          "3",
		Text:        "Design a simple web server in Go.",
		Hint:        "Look into the net/http package.",
		Constraints: "Must handle at least basic GET requests.",
		Solution:    "Use http.HandleFunc and http.ListenAndServe.",
		IsProject:   true,
	},
}

// getProblems returns a list of all the problems as JSON
// gin.context is important to gin. Mainly, t contains the request and serialize/deserialize the JSON
func getProblems(c *gin.Context) {
	//Context.IndentedJSON serializes the struct into JSON and adds it to the response
	//also returning HTTP_200 via http package and 'StatusOK', and then the data to be serialized
	c.IndentedJSON(http.StatusOK, problems)
}

// adds a new problem object from JSON POST data
// c is a pointer to an instance of 'gin.Context' (not a copy of it)
func postProblems(c *gin.Context) {
	var newProblem problem
	//BindJSON 'binds' the retrieved JSON body to newProblem (deserializing)
	//will return 400 if error occurs
	//the "&" passes a pointer to BindJSON, which is required
	if err := c.BindJSON(&newProblem); err != nil { //nil, kinda like 2-0, 'two - nil'
		return
	}
	//Add the new problem to the problems array
	problems = append(problems, newProblem)
	//return HTTP_201 along with the instance the user just created.
	c.IndentedJSON(http.StatusCreated, newProblem)
}

func getProblemByID(c *gin.Context) {
	requestID := c.Param("id")

	//loop over list of items to get the desired one
	for _, problem := range problems {
		if problem.ID == requestID {
			c.IndentedJSON(http.StatusOK, problem)
			return
		}
	}
	//if nothing
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "problem not found"})
}

func getIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	//initialize Gin Router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndexPage)
	router.GET("/problems", getProblems)
	router.POST("/problems", postProblems)
	router.GET("/problems/:id", getProblemByID)

	//attach router to http.Server and start it aswell
	router.Run("localhost:8080")
}
