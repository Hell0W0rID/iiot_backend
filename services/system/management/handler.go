package management

import (
	"net/http"

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

func (h *Handler) GetServices(c echo.Context) error {
	services, err := h.service.GetAllServices()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve services", err)
	}
	return utils.SuccessResponse(c, services)
}

func (h *Handler) GetService(c echo.Context) error {
	serviceID := c.Param("id")
	service, err := h.service.GetServiceByID(serviceID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Service not found", err)
	}
	return utils.SuccessResponse(c, service)
}

func (h *Handler) GetServiceByName(c echo.Context) error {
	name := c.Param("name")
	service, err := h.service.GetServiceByName(name)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Service not found", err)
	}
	return utils.SuccessResponse(c, service)
}

func (h *Handler) RegisterService(c echo.Context) error {
	var serviceInfo models.ServiceInfo
	if err := c.Bind(&serviceInfo); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&serviceInfo); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.RegisterService(&serviceInfo)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to register service", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) DeregisterService(c echo.Context) error {
	serviceID := c.Param("id")
	if err := h.service.DeregisterService(serviceID); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to deregister service", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Service deregistered successfully"})
}

func (h *Handler) UpdateServiceStatus(c echo.Context) error {
	serviceID := c.Param("id")
	var req struct {
		Status string `json:"status" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateServiceStatus(serviceID, req.Status); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update service status", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Service status updated successfully"})
}

func (h *Handler) GetServiceHealth(c echo.Context) error {
	serviceID := c.Param("id")
	health, err := h.service.GetServiceHealth(serviceID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Service health not found", err)
	}
	return utils.SuccessResponse(c, health)
}

func (h *Handler) GetAllServiceHealth(c echo.Context) error {
	healthStatuses, err := h.service.GetAllServiceHealth()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve service health", err)
	}
	return utils.SuccessResponse(c, healthStatuses)
}

func (h *Handler) PerformServiceOperation(c echo.Context) error {
	var req models.ServiceOperation
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	result, err := h.service.PerformOperation(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to perform operation", err)
	}

	return utils.SuccessResponse(c, result)
}

func (h *Handler) GetServiceConfig(c echo.Context) error {
	serviceID := c.Param("id")
	config, err := h.service.GetServiceConfig(serviceID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Service configuration not found", err)
	}
	return utils.SuccessResponse(c, config)
}

func (h *Handler) UpdateServiceConfig(c echo.Context) error {
	serviceID := c.Param("id")
	var config map[string]interface{}
	if err := c.Bind(&config); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateServiceConfig(serviceID, config); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update service configuration", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Service configuration updated successfully"})
}

func (h *Handler) GetServiceMetrics(c echo.Context) error {
	serviceID := c.Param("id")
	metrics, err := h.service.GetServiceMetrics(serviceID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Service metrics not found", err)
	}
	return utils.SuccessResponse(c, metrics)
}

func (h *Handler) GetSystemInfo(c echo.Context) error {
	systemInfo, err := h.service.GetSystemInfo()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve system information", err)
	}
	return utils.SuccessResponse(c, systemInfo)
}

func (h *Handler) GetSystemMetrics(c echo.Context) error {
	metrics, err := h.service.GetSystemMetrics()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve system metrics", err)
	}
	return utils.SuccessResponse(c, metrics)
}
