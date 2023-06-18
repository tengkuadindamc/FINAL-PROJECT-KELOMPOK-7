package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var todos []Todo
var currentID int

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	todo.ID = getNextID()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}

func getTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Todo not found"})
}

func updateTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			updatedTodo.ID = todo.ID
			updatedTodo.CreatedAt = todo.CreatedAt
			updatedTodo.UpdatedAt = time.Now()
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Todo not found"})
}

func deleteTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Todo not found"})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func getNextID() int {
	currentID++
	return currentID
}

func main() {
	r := gin.Default()

	r.GET("/swagger.json", func(c *gin.Context) {
		c.File("./swagger.json")
	})

	r.StaticFS("/swagger", http.Dir("swagger-ui"))

	v1 := r.Group("/api/v1")
	{
		todosGroup := v1.Group("/todos")
		{
			todosGroup.GET("", getAllTodos)
			todosGroup.POST("", createTodo)
			todosGroup.GET("/:id", getTodoByID)
			todosGroup.PUT("/:id", updateTodoByID)
			todosGroup.DELETE("/:id", deleteTodoByID)
		}
	}

	todos = []Todo{
		{ID: getNextID(), Title: "Task 1", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: getNextID(), Title: "Task 2", Completed: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
