package registry

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, service *Service) {
	handler := NewHandler(service)

	// Service registry routes
	registry := g.Group("/registry")
	registry.POST("/service", handler.RegisterService)
	registry.DELETE("/service/:id", handler.DeregisterService)
	registry.GET("/service", handler.GetServices)
	registry.GET("/service/:id", handler.GetService)
	registry.GET("/service/name/:name", handler.GetServiceByName)

	// Service discovery routes
	registry.GET("/discover", handler.DiscoverServices)
	registry.GET("/healthy", handler.GetHealthyServices)

	// Service health routes
	registry.PUT("/service/:id/health", handler.UpdateServiceHealth)
	registry.GET("/service/:id/health", handler.GetServiceHealth)

	// Service configuration routes
	registry.GET("/service/:id/config", handler.GetServiceConfig)
	registry.PUT("/service/:id/config", handler.UpdateServiceConfig)
}
