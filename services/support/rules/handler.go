package rules

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

// Rule handlers
func (h *Handler) GetRules(c echo.Context) error {
	enabled := c.QueryParam("enabled")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	var enabledFilter *bool
	if enabled != "" {
		enabledValue := enabled == "true"
		enabledFilter = &enabledValue
	}

	rules, err := h.service.GetRules(enabledFilter, limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve rules", err)
	}
	return utils.SuccessResponse(c, rules)
}

func (h *Handler) GetRule(c echo.Context) error {
	id := c.Param("id")
	rule, err := h.service.GetRuleByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Rule not found", err)
	}
	return utils.SuccessResponse(c, rule)
}

func (h *Handler) GetRuleByName(c echo.Context) error {
	name := c.Param("name")
	rule, err := h.service.GetRuleByName(name)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Rule not found", err)
	}
	return utils.SuccessResponse(c, rule)
}

func (h *Handler) CreateRule(c echo.Context) error {
	var req models.RuleRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.CreateRule(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create rule", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) UpdateRule(c echo.Context) error {
	id := c.Param("id")
	var req models.RuleRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdateRule(id, &req); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update rule", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Rule updated successfully"})
}

func (h *Handler) DeleteRule(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteRule(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete rule", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Rule deleted successfully"})
}

func (h *Handler) EnableRule(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.SetRuleEnabled(id, true); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to enable rule", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Rule enabled successfully"})
}

func (h *Handler) DisableRule(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.SetRuleEnabled(id, false); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to disable rule", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Rule disabled successfully"})
}

// Pipeline handlers
func (h *Handler) GetPipelines(c echo.Context) error {
	enabled := c.QueryParam("enabled")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	var enabledFilter *bool
	if enabled != "" {
		enabledValue := enabled == "true"
		enabledFilter = &enabledValue
	}

	pipelines, err := h.service.GetPipelines(enabledFilter, limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve pipelines", err)
	}
	return utils.SuccessResponse(c, pipelines)
}

func (h *Handler) GetPipeline(c echo.Context) error {
	id := c.Param("id")
	pipeline, err := h.service.GetPipelineByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Pipeline not found", err)
	}
	return utils.SuccessResponse(c, pipeline)
}

func (h *Handler) GetPipelineByName(c echo.Context) error {
	name := c.Param("name")
	pipeline, err := h.service.GetPipelineByName(name)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Pipeline not found", err)
	}
	return utils.SuccessResponse(c, pipeline)
}

func (h *Handler) CreatePipeline(c echo.Context) error {
	var req models.PipelineRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err)
	}

	id, err := h.service.CreatePipeline(&req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create pipeline", err)
	}

	return utils.SuccessResponse(c, map[string]string{"id": id})
}

func (h *Handler) UpdatePipeline(c echo.Context) error {
	id := c.Param("id")
	var req models.PipelineRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.service.UpdatePipeline(id, &req); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update pipeline", err)
	}

	return utils.SuccessResponse(c, map[string]string{"message": "Pipeline updated successfully"})
}

func (h *Handler) DeletePipeline(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeletePipeline(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete pipeline", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Pipeline deleted successfully"})
}

func (h *Handler) EnablePipeline(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.SetPipelineEnabled(id, true); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to enable pipeline", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Pipeline enabled successfully"})
}

func (h *Handler) DisablePipeline(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.SetPipelineEnabled(id, false); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to disable pipeline", err)
	}
	return utils.SuccessResponse(c, map[string]string{"message": "Pipeline disabled successfully"})
}

func (h *Handler) ExecuteRule(c echo.Context) error {
	id := c.Param("id")
	execution, err := h.service.ExecuteRule(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to execute rule", err)
	}
	return utils.SuccessResponse(c, execution)
}

func (h *Handler) GetRuleExecutions(c echo.Context) error {
	ruleID := c.QueryParam("ruleId")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit == 0 {
		limit = 50
	}

	executions, err := h.service.GetRuleExecutions(ruleID, limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve rule executions", err)
	}
	return utils.SuccessResponse(c, executions)
}
