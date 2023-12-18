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
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func main() {
	e := echo.New()
	e.GET("/", contentHandler)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":8080"))
}

// Code from:
// https://echo.labstack.com/docs/quick-start#path-parameters
