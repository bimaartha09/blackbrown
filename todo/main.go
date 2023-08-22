package main

import (
	"main/todo/method"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", method.GetTodos)
	router.POST("/todos", method.AddTodos)
	router.GET("/todos/:id", method.GetTodo)
	router.Run("localhost:9090")
}
