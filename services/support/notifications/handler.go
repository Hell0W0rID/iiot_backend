package notifications

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"iiot-backend/models"
	"iiot-backend/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Notification handlers
func (h *Handler) GetNotifications(c echo.Context) error {
	category := c.QueryParam("category")
	severity := c.QueryParam("severity")
	status := c.QueryParam("status")
	startStr := c.QueryParam("start")
	endStr := c.QueryParam("end")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	filter := models.NotificationFilter{
		Category: category,
		Severity: severity,
		Status:   status,
		Limit:    limit,
		Offset:   offset,
	}

	// Parse start and end times
	if startStr != "" {
		if start, err := time.Parse(time.RFC3339, startStr); err == nil {
			filter.Start = start
		}
	}
	if endStr != "" {
		if end, err := time.Parse(time.RFC3339, endStr); err == nil {
			filter.End = end
		}
	}

	notifications, err := h.service.GetNotifications(filter)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve notifications", err)
	}

	return utils.SuccessResponse(c, notifications)
}

func (h *Handler) GetNotification(c echo.Context) error {
	id := c.Param("id")
	notification, err := h.service.GetNotificationByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Notification not found", err)
	}
	return utils.SuccessResponse(c, notification)
}

func (h *Handler) GetNotificationBySlug(c echo.Context) error {
	slug := c.Param("slug")
	notification, err := h.service.GetNotificationBySlug(slug)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Notification not found", err)
	}
	return utils.SuccessResponse(c, notification)
}

func (h *Handler) CreateNotification(c echo.Context) error {
	var req models.NotificationRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.CreateNotification(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create notification", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) UpdateNotification(c echo.Context) error {
	id := c.Param("id")
	var req models.NotificationRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateNotification(id, &req); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update notification", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Notification updated successfully"})
}

func (h *Handler) DeleteNotification(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteNotification(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete notification", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Notification deleted successfully"})
}

func (h *Handler) DeleteNotificationBySlug(c echo.Context) error {
	slug := c.Param("slug")
	if err := h.service.DeleteNotificationBySlug(slug); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete notification", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Notification deleted successfully"})
}

func (h *Handler) CleanupNotifications(c echo.Context) error {
	ageStr := c.Param("age")
	age, err := strconv.ParseInt(ageStr, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid age parameter", err)
	}

	count, err := h.service.CleanupNotifications(age)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to cleanup notifications", err)
	}

	return utils.SuccessResponse(c, map[string]interface{}{
		"message": "Notifications cleaned up successfully",
		"count":   count,
	})
}

// Subscription handlers
func (h *Handler) GetSubscriptions(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	subscriptions, err := h.service.GetAllSubscriptions(limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve subscriptions", err)
	}
	return utils.SuccessResponse(c, subscriptions)
}

func (h *Handler) GetSubscription(c echo.Context) error {
	id := c.Param("id")
	subscription, err := h.service.GetSubscriptionByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Subscription not found", err)
	}
	return utils.SuccessResponse(c, subscription)
}

func (h *Handler) GetSubscriptionByName(c echo.Context) error {
	name := c.Param("name")
	subscription, err := h.service.GetSubscriptionByName(name)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Subscription not found", err)
	}
	return utils.SuccessResponse(c, subscription)
}

func (h *Handler) CreateSubscription(c echo.Context) error {
	var req models.SubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.CreateSubscription(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create subscription", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) UpdateSubscription(c echo.Context) error {
	id := c.Param("id")
	var req models.SubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateSubscription(id, &req); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update subscription", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Subscription updated successfully"})
}

func (h *Handler) DeleteSubscription(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteSubscription(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete subscription", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Subscription deleted successfully"})
}

func (h *Handler) TransmitNotification(c echo.Context) error {
	id := c.Param("id")
	err := h.service.TransmitNotification(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to transmit notification", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Notification transmitted successfully"})
}
