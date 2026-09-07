package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	echoSwagger "github.com/swaggo/echo-swagger/v2"

	_ "github.com/maruf03/markme/backend/docs"
)

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/livez", livez)
	e.GET("/healthz", healthz)
	e.GET("/readyz", readyz)
	e.GET("/docs/*", echoSwagger.WrapHandler)

	var v1_group = e.Group("/api/v1")
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
