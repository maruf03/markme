package main

import (
	"github.com/labstack/echo/v5"
)

// @title			MarkMe API
// @version		1.0
// @description	This is a sample server Petstore server.
func main() {
	e := echo.New()

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
