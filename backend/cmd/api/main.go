package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	echootel "github.com/labstack/echo-opentelemetry"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	authController "github.com/maruf03/markme/backend/cmd/api/auth"
	bookmarkController "github.com/maruf03/markme/backend/cmd/api/bookmark"
	userController "github.com/maruf03/markme/backend/cmd/api/user"
	_ "github.com/maruf03/markme/backend/docs"
	"github.com/maruf03/markme/backend/internal/auth"
	"github.com/maruf03/markme/backend/internal/bookmark"
	"github.com/maruf03/markme/backend/internal/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

// @title			MarkMe API
// @version		1.0
// @description	This is a sample server Petstore server.
func main() {
	ctx := context.Background()

	// 1. Create the Prometheus OTel exporter
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("failed to initialize prometheus exporter: %v", err)
	}

	// 2. Set up the MeterProvider with the Prometheus reader
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	otel.SetMeterProvider(provider)
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			log.Printf("Error terminating MeterProvider: %v", err)
		}
	}()

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	RegisterMiddlewares(e)
	RegisterRoutes(e)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(echootel.NewMiddleware("app.example.com"))
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/livez", livez)
	e.GET("/healthz", healthz)
	e.GET("/readyz", readyz)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/docs/*", echoSwagger.WrapHandler)

	v1_group := e.Group("/api/v1")

	newAuthController := authController.NewAuthController(
		&auth.AuthService{},
		v1_group.Group("/auth"),
	)
	newAuthController.RegisterRoutes()

	newUserController := userController.NewUserController(
		&user.UserService{},
		v1_group.Group("/user"),
	)
	newUserController.RegisterRoutes()

	newBookmarkController := bookmarkController.NewBookmarkController(
		&bookmark.BookmarkService{},
		v1_group.Group("/user/:userId/bookmark"),
	)
	newBookmarkController.RegisterRoutes()
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
