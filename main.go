package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func getTodos(c echo.Context) {
    c.String(http.StatusOK, "Get Todos")
}

func addTodo(c echo.Context) {
    c.String(http.StatusOK, "Add Todo")
}

func removeTodo(c echo.Context) {
    c.String(http.StatusOK, "Remove Todo")
}

func updateTodo(c echo.Context) {
    c.String(http.StatusOK, "Update Todo")
}

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":9090"))
}