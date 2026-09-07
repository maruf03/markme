package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/maruf03/markme/backend/internal/auth"
)

type AuthController struct {
	authService *auth.AuthService
	routeGroup  *echo.Group
}

func NewAuthController(authService *auth.AuthService, routeGroup *echo.Group) *AuthController {
	return &AuthController{
		authService: authService,
		routeGroup:  routeGroup,
	}
}

func (c *AuthController) RegisterRoutes() {
	c.routeGroup.POST("/login", c.Login)
	c.routeGroup.POST("/refresh", c.Refresh)
}

func (c *AuthController) Login(e *echo.Context) error {
	request := new(LoginRequest)
	if err := e.Bind(request); err != nil {
		return err
	}
	if err := e.Validate(request); err != nil {
		return err
	}
	response, err := c.authService.Login(request.Email, request.Password)
	if err != nil {
		return err
	}
	return e.JSON(
		http.StatusOK,
		LoginResponse{AccessToken: response.AccessToken, RefreshToken: response.RefreshToken},
	)
}

func (c *AuthController) Refresh(e *echo.Context) error {
	request := new(RefreshRequest)
	if err := e.Bind(request); err != nil {
		return err
	}
	if err := e.Validate(request); err != nil {
		return err
	}
	response, err := c.authService.Refresh(request.RefreshToken)
	if err != nil {
		return err
	}
	return e.JSON(
		http.StatusOK,
		LoginResponse{AccessToken: response.AccessToken, RefreshToken: response.RefreshToken},
	)
}
