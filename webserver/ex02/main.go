package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func contentHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func main() {
	e := echo.New()
	e.GET("/", contentHandler)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.Logger.Fatal(e.Start(":8080"))
}

// Code from:
// https://echo.labstack.com/docs/quick-start#form-applicationx-www-form-urlencoded
