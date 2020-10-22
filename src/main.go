package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
	"time"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	t := &Template{
		templates: template.Must(template.ParseGlob("*.html")),
	}

	e.Renderer = t

	e.GET("", index)
	e.POST("/login", login)
	e.GET("/as", doneTodo)

	// Restricted group
	r := e.Group("/a")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", index)
	r.GET("aa", index)
	r.POST("add", addTodo)
	r.POST("delete", deleteToDo)
	r.POST("done", doneTodo)

	e.Logger.Fatal(e.Start(":2020"))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)

}

func addTodo(c echo.Context) error {

	return c.String(http.StatusOK, "ADD \n")
}

func deleteToDo(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func doneTodo(c echo.Context) error {

	return c.String(http.StatusOK, "SAFASDFAS")
}


func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	println(username)
	println(password)

	// Throws unauthorized error
	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}


