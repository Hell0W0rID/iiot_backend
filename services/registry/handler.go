package registry

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

func (h *Handler) UpdateServiceHealth(c echo.Context) error {
	serviceID := c.Param("id")
	var healthUpdate struct {
		Status  string `json:"status" validate:"required"`
		Message string `json:"message"`
	}

	if err := c.Bind(&healthUpdate); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateServiceHealth(serviceID, healthUpdate.Status, healthUpdate.Message); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update service health", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Service health updated successfully"})
}

func (h *Handler) GetServiceHealth(c echo.Context) error {
	serviceID := c.Param("id")
	health, err := h.service.GetServiceHealth(serviceID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Service health not found", err)
	}
	return utils.SuccessResponse(c, health)
}

func (h *Handler) GetHealthyServices(c echo.Context) error {
	services, err := h.service.GetHealthyServices()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve healthy services", err)
	}
	return utils.SuccessResponse(c, services)
}

func (h *Handler) DiscoverServices(c echo.Context) error {
	serviceType := c.QueryParam("type")
	tag := c.QueryParam("tag")
	
	services, err := h.service.DiscoverServices(serviceType, tag)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to discover services", err)
	}
	return utils.SuccessResponse(c, services)
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
