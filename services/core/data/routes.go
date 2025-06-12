package data

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, service *Service) {
	handler := NewHandler(service)

	// Events routes
	events := g.Group("/event")
	events.GET("", handler.GetEvents)
	events.GET("/:id", handler.GetEvent)
	events.POST("", handler.CreateEvent)
	events.DELETE("/:id", handler.DeleteEvent)
	events.DELETE("/device/:device", handler.DeleteEventsByDevice)
	events.DELETE("/age/:age", handler.DeleteEventsByAge)

	// Readings routes
	readings := g.Group("/reading")
	readings.GET("", handler.GetReadings)
	readings.GET("/:id", handler.GetReading)
	readings.DELETE("/:id", handler.DeleteReading)

	// Count routes
	g.GET("/event/count", handler.GetEventCount)
	g.GET("/event/count/device/:device", handler.GetEventCountByDevice)
	g.GET("/reading/count", handler.GetReadingCount)
	g.GET("/reading/count/device/:device", handler.GetReadingCountByDevice)
}
