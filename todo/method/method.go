package method

import (
	"errors"
	"main/todo/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, entity.Todos)
}

func AddTodos(context *gin.Context) {
	var newTodo entity.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, newTodo)
		return
	}

	if newTodo.ID == "" || newTodo.Item == "" {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": "Required param not to be empty"})
		return
	}

	for _, todo := range entity.Todos {
		if todo.ID == newTodo.ID {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": "ID are same with previous data"})
			return
		}
	}

	entity.Todos = append(entity.Todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(id string) (*entity.Todo, error) {
	for _, todo := range entity.Todos {
		if todo.ID == id {
			return &todo, nil
		}
	}

	return nil, errors.New("ID not found")
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"errors": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusFound, todo)
}
