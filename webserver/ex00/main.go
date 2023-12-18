package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func contentHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", contentHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

// Code from:
// https://echo.labstack.com/docs/quick-start#handling-request
