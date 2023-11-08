package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
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
    return c.HTML(http.StatusOK, "<h1>Add Todo</h1>")
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
    
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://127.0.0.1:5500"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Hx-Current-Url", "Hx-Request",},
    }))

    e.GET("/todos", getTodos)
    e.POST("/addTodo", addTodo)
    e.DELETE("/removeTodo/:id", removeTodo)
    e.PATCH("/updateTodo/:id", updateTodo)
    e.Logger.Fatal(e.Start(":9090"))
}