package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	authController "github.com/maruf03/markme/backend/cmd/api/auth"
	bookmarkController "github.com/maruf03/markme/backend/cmd/api/bookmark"
	userController "github.com/maruf03/markme/backend/cmd/api/user"
	_ "github.com/maruf03/markme/backend/docs"
	"github.com/maruf03/markme/backend/internal/auth"
	"github.com/maruf03/markme/backend/internal/bookmark"
	"github.com/maruf03/markme/backend/internal/user"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
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
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/livez", livez)
	e.GET("/healthz", healthz)
	e.GET("/readyz", readyz)
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
		v1_group.Group("/bookmark"),
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
