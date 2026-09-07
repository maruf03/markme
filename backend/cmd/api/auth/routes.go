package auth

import "github.com/labstack/echo/v5"

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) RegisterRoutes(g *echo.Group) {

}

func (c *AuthController) Login() {}
