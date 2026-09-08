package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	echootel "github.com/labstack/echo-opentelemetry"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	authController "github.com/maruf03/markme/backend/cmd/api/auth"
	bookmarkController "github.com/maruf03/markme/backend/cmd/api/bookmark"
	obsController "github.com/maruf03/markme/backend/cmd/api/observability"
	userController "github.com/maruf03/markme/backend/cmd/api/user"
	_ "github.com/maruf03/markme/backend/docs"
	"github.com/maruf03/markme/backend/internal/auth"
	"github.com/maruf03/markme/backend/internal/bookmark"
	"github.com/maruf03/markme/backend/internal/observability"
	"github.com/maruf03/markme/backend/internal/user"
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

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
	e.Use(echootel.NewMiddleware("app.example.com")) // TODO: replace with real value
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/docs/*", echoSwagger.WrapHandler)

	obs_group := e.Group("")

	newObsController := obsController.NewObservabilityController(
		&observability.ObservabilityService{},
		obs_group,
	)

	newObsController.RegisterRoutes()

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
