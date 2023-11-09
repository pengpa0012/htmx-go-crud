package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "math/rand"
    "time"
    "strconv"
)

type Todo struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Description  string  `json:"description"`
	DateCreated  time.Time  `json:"date_created"`
	Completed  bool `json:"completed"`
}

var todos = []Todo {
    {ID: 1, Title: "Work", Description: "Do the thing", DateCreated: time.Now(), Completed: false},
	{ID: 2, Title: "Personal Time", Description: "Go to gym", DateCreated: time.Now(), Completed: false},
	{ID: 3, Title: "Fun Time", Description: "Play games", DateCreated: time.Now(), Completed: false},
}

func getTodos(c echo.Context) error {
    return c.JSON(http.StatusOK, todos)
}

func addTodo(c echo.Context) error {
    rand.Seed(time.Now().UnixNano())
    title := c.FormValue("title")
    description := c.FormValue("description")
    todo := Todo {
        ID: rand.Intn(1000), 
        Title: title,
        Description: description,
        DateCreated: time.Now(),
    }

    todos = append(todos, todo)

    return c.JSON(http.StatusOK, todos)
}

func removeTodo(c echo.Context) error {
    id := c.Param("id")

    parseID, _ := strconv.Atoi(id)

    for index, todo := range todos {
        if todo.ID == parseID {
            todos = append(todos[:index], todos[index+1:]...)
            return c.JSON(http.StatusCreated, todos) 
        }
    }

    return c.JSON(http.StatusNotFound, "Error remove todo")
}

func updateTodo(c echo.Context) error {
    id := c.Param("id")
    title := c.FormValue("title")
    description := c.FormValue("description")
    parseID, _ := strconv.Atoi(id)

    for index, todo := range todos {
        if todo.ID == parseID {
            todos[index].Title = title
            todos[index].Description = description
            return c.JSON(http.StatusCreated, todos) 
        }
    }

    return c.JSON(http.StatusOK, id)
}

func main() {
    e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://127.0.0.1:5500"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Hx-Current-Url", "Hx-Request", "Hx-Target"},
    }))

    e.GET("/todos", getTodos)
    e.POST("/addTodo", addTodo)
    e.DELETE("/removeTodo/:id", removeTodo)
    e.PATCH("/updateTodo/:id", updateTodo)
    e.Logger.Fatal(e.Start(":9090"))
}