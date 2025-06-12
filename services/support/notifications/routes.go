package notifications

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, service *Service) {
	handler := NewHandler(service)

	// Notifications routes
	notifications := g.Group("/notification")
	notifications.GET("", handler.GetNotifications)
	notifications.GET("/:id", handler.GetNotification)
	notifications.GET("/slug/:slug", handler.GetNotificationBySlug)
	notifications.POST("", handler.CreateNotification)
	notifications.PUT("/:id", handler.UpdateNotification)
	notifications.DELETE("/:id", handler.DeleteNotification)
	notifications.DELETE("/slug/:slug", handler.DeleteNotificationBySlug)
	notifications.DELETE("/age/:age", handler.CleanupNotifications)
	notifications.POST("/:id/transmission", handler.TransmitNotification)

	// Subscriptions routes
	subscriptions := g.Group("/subscription")
	subscriptions.GET("", handler.GetSubscriptions)
	subscriptions.GET("/:id", handler.GetSubscription)
	subscriptions.GET("/name/:name", handler.GetSubscriptionByName)
	subscriptions.POST("", handler.CreateSubscription)
	subscriptions.PUT("/:id", handler.UpdateSubscription)
	subscriptions.DELETE("/:id", handler.DeleteSubscription)
}
