package scheduler

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, service *Service) {
	handler := NewHandler(service)

	// Intervals routes
	intervals := g.Group("/interval")
	intervals.GET("", handler.GetIntervals)
	intervals.GET("/:id", handler.GetInterval)
	intervals.GET("/name/:name", handler.GetIntervalByName)
	intervals.POST("", handler.CreateInterval)
	intervals.PUT("/:id", handler.UpdateInterval)
	intervals.DELETE("/:id", handler.DeleteInterval)

	// Interval Actions routes
	intervalActions := g.Group("/intervalaction")
	intervalActions.GET("", handler.GetIntervalActions)
	intervalActions.GET("/:id", handler.GetIntervalAction)
	intervalActions.GET("/name/:name", handler.GetIntervalActionByName)
	intervalActions.POST("", handler.CreateIntervalAction)
	intervalActions.PUT("/:id", handler.UpdateIntervalAction)
	intervalActions.DELETE("/:id", handler.DeleteIntervalAction)

	// Schedule Status routes
	g.GET("/schedule/status/:name", handler.GetScheduleStatus)
	g.GET("/schedule/status", handler.GetAllScheduleStatuses)
}
