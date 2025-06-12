package scheduler

import (
	"net/http"
	"strconv"

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

// Interval handlers
func (h *Handler) GetIntervals(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	intervals, err := h.service.GetAllIntervals(limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve intervals", err)
	}
	return utils.SuccessResponse(c, intervals)
}

func (h *Handler) GetInterval(c echo.Context) error {
	id := c.Param("id")
	interval, err := h.service.GetIntervalByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Interval not found", err)
	}
	return utils.SuccessResponse(c, interval)
}

func (h *Handler) GetIntervalByName(c echo.Context) error {
	name := c.Param("name")
	interval, err := h.service.GetIntervalByName(name)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Interval not found", err)
	}
	return utils.SuccessResponse(c, interval)
}

func (h *Handler) CreateInterval(c echo.Context) error {
	var req models.IntervalRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.CreateInterval(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create interval", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) UpdateInterval(c echo.Context) error {
	id := c.Param("id")
	var req models.IntervalRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateInterval(id, &req); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update interval", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Interval updated successfully"})
}

func (h *Handler) DeleteInterval(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteInterval(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete interval", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Interval deleted successfully"})
}

// Interval Action handlers
func (h *Handler) GetIntervalActions(c echo.Context) error {
	intervalName := c.QueryParam("intervalName")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	actions, err := h.service.GetIntervalActions(intervalName, limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve interval actions", err)
	}
	return utils.SuccessResponse(c, actions)
}

func (h *Handler) GetIntervalAction(c echo.Context) error {
	id := c.Param("id")
	action, err := h.service.GetIntervalActionByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Interval action not found", err)
	}
	return utils.SuccessResponse(c, action)
}

func (h *Handler) GetIntervalActionByName(c echo.Context) error {
	name := c.Param("name")
	action, err := h.service.GetIntervalActionByName(name)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Interval action not found", err)
	}
	return utils.SuccessResponse(c, action)
}

func (h *Handler) CreateIntervalAction(c echo.Context) error {
	var req models.IntervalActionRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.CreateIntervalAction(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create interval action", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) UpdateIntervalAction(c echo.Context) error {
	id := c.Param("id")
	var req models.IntervalActionRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateIntervalAction(id, &req); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update interval action", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Interval action updated successfully"})
}

func (h *Handler) DeleteIntervalAction(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteIntervalAction(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete interval action", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Interval action deleted successfully"})
}

// Schedule Status handlers
func (h *Handler) GetScheduleStatus(c echo.Context) error {
	intervalName := c.Param("name")
	status, err := h.service.GetScheduleStatus(intervalName)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Schedule status not found", err)
	}
	return utils.SuccessResponse(c, status)
}

func (h *Handler) GetAllScheduleStatuses(c echo.Context) error {
	statuses, err := h.service.GetAllScheduleStatuses()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve schedule statuses", err)
	}
	return utils.SuccessResponse(c, statuses)
}
