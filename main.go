package main

import (
    "io"
    "net/http"
    "github.com/labstack/echo/v4"
    "math/rand"
    "time"
    "strconv"
    "html/template"
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

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func getTodos(c echo.Context) error {
    return c.Render(http.StatusOK, "todos.html", todos)
}

func addTodo(c echo.Context) error {
    rand.Seed(time.Now().UnixNano())
    title := c.FormValue("title")
    description := c.FormValue("description")

    if title == "" || description == "" {
        return c.Render(http.StatusOK, "todos.html", todos)
    }

    todo := Todo {
        ID: rand.Intn(1000), 
        Title: title,
        Description: description,
        DateCreated: time.Now(),
    }

    todos = append(todos, todo)

    return c.Render(http.StatusOK, "todos.html", todos)
}

func removeTodo(c echo.Context) error {
    id := c.Param("id")

    parseID, _ := strconv.Atoi(id)

    for index, todo := range todos {
        if todo.ID == parseID {
            todos = append(todos[:index], todos[index+1:]...)
            return c.Render(http.StatusOK, "todos.html", todos)
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
            return c.Render(http.StatusOK, "todos.html", todos)
        }
    }

    return c.JSON(http.StatusOK, id)
}

func Home(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", "/")
}

func main() {
    e := echo.New()

    t := &Template{
        templates: template.Must(template.ParseGlob("web/templates/*.html")),
    }

    e.Renderer = t
    e.GET("/", Home)
    e.GET("/todos", getTodos)
    e.POST("/addTodo", addTodo)
    e.DELETE("/removeTodo/:id", removeTodo)
    e.PATCH("/updateTodo/:id", updateTodo)
    e.Logger.Fatal(e.Start(":5000"))
}