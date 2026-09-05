package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	echoSwagger "github.com/swaggo/echo-swagger/v2"

	_ "github.com/maruf03/markme/backend/docs"
)

// @title			MarkMe API
// @version		1.0
// @description	This is a sample server Petstore server.
func main() {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/livez", livez)
	e.GET("/healthz", healthz)
	e.GET("/readyz", readyz)

	e.GET("/docs/*", echoSwagger.WrapHandler)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func livez(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func healthz(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func readyz(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
