package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Todo struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Description  string  `json:"description"`
	DateCreated  string  `json:"date_created"`
	Completed  bool `json:"completed"`
}

func getTodos(c echo.Context) error {
    return c.String(http.StatusOK, "Get Todos")
}

func addTodo(c echo.Context) error {
    return c.String(http.StatusOK, "Add Todo")
}

func removeTodo(c echo.Context) error {
    return c.String(http.StatusOK, "Remove Todo")
}

func updateTodo(c echo.Context) error {
    return c.String(http.StatusOK, "Update Todo")
}

func main() {
    e := echo.New()
    e.GET("/todos", getTodos)
    e.POST("/addTodo", addTodo)
    e.DELETE("/removeTodo/:id", removeTodo)
    e.PATCH("/updateTodo/:id", updateTodo)
    e.Logger.Fatal(e.Start(":9090"))
}