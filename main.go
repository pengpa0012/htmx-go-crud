package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Todo struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Description  string  `json:"description"`
	DateCreated  int  `json:"date_created"`
	Completed  bool `json:"completed"`
}

var todos = []Todo {
    {ID: "1", Title: "Work", Description: "Do the thing", DateCreated: 1699404506241, Completed: false},
	{ID: "2", Title: "Personal Time", Description: "Go to gym", DateCreated: 1699404506241, Completed: false},
	{ID: "3", Title: "Fun Time", Description: "Play games", DateCreated: 1699404506241, Completed: false},
}

func getTodos(c echo.Context) error {
    return c.JSON(http.StatusOK, todos)
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