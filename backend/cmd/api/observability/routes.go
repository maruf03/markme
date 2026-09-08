package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	obs "github.com/maruf03/markme/backend/internal/observability"
	"github.com/maruf03/markme/backend/internal/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ObservabilityController struct {
	obsService *obs.ObservabilityService
	routeGroup *echo.Group
}

func NewObservabilityController(
	obsService *obs.ObservabilityService,
	routeGroup *echo.Group,
) *ObservabilityController {
	return &ObservabilityController{
		obsService: obsService,
		routeGroup: routeGroup,
	}
}

func (o *ObservabilityController) RegisterRoutes() {
	o.routeGroup.GET("/livez", o.livez)
	o.routeGroup.GET("/healthz", o.healthz)
	o.routeGroup.GET("/readyz", o.readyz)
	o.routeGroup.GET("/version", o.version)
	o.routeGroup.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}

func (o *ObservabilityController) livez(c *echo.Context) error {
	return c.JSON(http.StatusOK, LivezResponse{Status: Pass})
}

func (o *ObservabilityController) healthz(c *echo.Context) error {
	return c.JSON(http.StatusOK, HealthzResponse{Status: Pass, Checks: nil})
}

func (o *ObservabilityController) readyz(c *echo.Context) error {
	return c.JSON(http.StatusOK, ReadyzResponse{Status: Pass, Checks: nil})
}

func (o *ObservabilityController) version(c *echo.Context) error {
	return c.JSON(http.StatusOK, VersionResponse{Version: version.Version, Commit: version.Commit, BuildTime: version.BuildTime})
}
