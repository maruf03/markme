package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/maruf03/markme/backend/internal/user"
)

type UserController struct {
	userService *user.UserService
	routeGroup  *echo.Group
}

func NewUserController(userService *user.UserService, routeGroup *echo.Group) *UserController {
	return &UserController{
		userService: userService,
		routeGroup:  routeGroup,
	}
}

func (c *UserController) RegisterRoutes() {
	c.routeGroup.GET("/me", c.GetProfile)
	c.routeGroup.PATCH("/me", c.UpdateProfile)
}

func (c *UserController) GetProfile(e *echo.Context) error {
	id := uuid.New()
	response, err := c.userService.GetProfile(id)
	if err != nil {
		return err
	}
	return e.JSON(
		http.StatusOK,
		ProfileResponse{
			Email:     response.Email,
			UserName:  response.UserName,
			FirstName: response.FirstName,
			LastName:  response.LastName,
			TagLine:   response.TagLine,
		},
	)
}

func (c *UserController) UpdateProfile(e *echo.Context) error {
	id := uuid.New()
	request := new(ProfileUpdateRequest)
	if err := e.Bind(request); err != nil {
		return err
	}
	if err := e.Validate(request); err != nil {
		return err
	}
	response, err := c.userService.UpdateProfile(id)
	if err != nil {
		return err
	}
	return e.JSON(
		http.StatusOK,
		ProfileResponse{
			Email:     response.Email,
			UserName:  response.UserName,
			FirstName: response.FirstName,
			LastName:  response.LastName,
			TagLine:   response.TagLine,
		},
	)
}
