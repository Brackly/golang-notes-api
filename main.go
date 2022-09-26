package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"Completed"`
}

var todos = []todo{
	{ID: "1", Item: "Read a book", Completed: false},
	{ID: "2", Item: "Talk to friends", Completed: false},
	{ID: "3", Item: "Take a walk", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func home(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "This is the home endpoint")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/", home)
	router.Run("localhost:3000")
}
