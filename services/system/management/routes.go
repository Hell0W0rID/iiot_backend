package management

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, service *Service) {
	handler := NewHandler(service)

	// Service management routes
	services := g.Group("/service")
	services.GET("", handler.GetServices)
	services.GET("/:id", handler.GetService)
	services.GET("/name/:name", handler.GetServiceByName)
	services.POST("", handler.RegisterService)
	services.DELETE("/:id", handler.DeregisterService)
	services.PUT("/:id/status", handler.UpdateServiceStatus)

	// Service operations
	services.POST("/operation", handler.PerformServiceOperation)

	// Service configuration
	services.GET("/:id/config", handler.GetServiceConfig)
	services.PUT("/:id/config", handler.UpdateServiceConfig)

	// Service health and metrics
	services.GET("/:id/health", handler.GetServiceHealth)
	services.GET("/:id/metrics", handler.GetServiceMetrics)

	// System-wide routes
	g.GET("/service/health", handler.GetAllServiceHealth)
	g.GET("/system/info", handler.GetSystemInfo)
	g.GET("/system/metrics", handler.GetSystemMetrics)
}
