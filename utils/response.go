package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// StandardResponse represents a standard API response
type StandardResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     *ErrorInfo  `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	RequestID string      `json:"request_id,omitempty"`
}

// ErrorInfo contains error details
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// ListResponse represents a paginated list response
type ListResponse struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	HasMore    bool        `json:"has_more"`
}

// SuccessResponse returns a successful response
func SuccessResponse(c echo.Context, data interface{}) error {
	response := StandardResponse{
		Success:   true,
		Data:      data,
		Timestamp: time.Now(),
		RequestID: getRequestID(c),
	}
	return c.JSON(http.StatusOK, response)
}

// SuccessResponseWithStatus returns a successful response with custom status code
func SuccessResponseWithStatus(c echo.Context, status int, data interface{}) error {
	response := StandardResponse{
		Success:   true,
		Data:      data,
		Timestamp: time.Now(),
		RequestID: getRequestID(c),
	}
	return c.JSON(status, response)
}

// ErrorResponse returns an error response
func ErrorResponse(c echo.Context, status int, message string, err error) error {
	errorInfo := &ErrorInfo{
		Code:    getErrorCode(status),
		Message: message,
	}

	if err != nil {
		errorInfo.Details = err.Error()
	}

	response := StandardResponse{
		Success:   false,
		Error:     errorInfo,
		Timestamp: time.Now(),
		RequestID: getRequestID(c),
	}

	return c.JSON(status, response)
}

// ValidationErrorResponse returns a validation error response
func ValidationErrorResponse(c echo.Context, validationErrors []ValidationError) error {
	errorInfo := &ErrorInfo{
		Code:    "VALIDATION_ERROR",
		Message: "Request validation failed",
		Details: formatValidationErrors(validationErrors),
	}

	response := StandardResponse{
		Success:   false,
		Error:     errorInfo,
		Timestamp: time.Now(),
		RequestID: getRequestID(c),
	}

	return c.JSON(http.StatusBadRequest, response)
}

// NotFoundResponse returns a not found error response
func NotFoundResponse(c echo.Context, resource string) error {
	return ErrorResponse(c, http.StatusNotFound, resource+" not found", nil)
}

// UnauthorizedResponse returns an unauthorized error response
func UnauthorizedResponse(c echo.Context, message string) error {
	if message == "" {
		message = "Unauthorized access"
	}
	return ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

// ForbiddenResponse returns a forbidden error response
func ForbiddenResponse(c echo.Context, message string) error {
	if message == "" {
		message = "Access forbidden"
	}
	return ErrorResponse(c, http.StatusForbidden, message, nil)
}

// ConflictResponse returns a conflict error response
func ConflictResponse(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusConflict, message, nil)
}

// InternalErrorResponse returns an internal server error response
func InternalErrorResponse(c echo.Context, err error) error {
	return ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err)
}

// ListSuccessResponse returns a successful paginated list response
func ListSuccessResponse(c echo.Context, items interface{}, totalCount int64, page, pageSize int) error {
	listData := ListResponse{
		Items:      items,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		HasMore:    int64((page-1)*pageSize+pageSize) < totalCount,
	}

	return SuccessResponse(c, listData)
}

// CreatedResponse returns a successful creation response
func CreatedResponse(c echo.Context, data interface{}) error {
	return SuccessResponseWithStatus(c, http.StatusCreated, data)
}

// AcceptedResponse returns an accepted response for async operations
func AcceptedResponse(c echo.Context, data interface{}) error {
	return SuccessResponseWithStatus(c, http.StatusAccepted, data)
}

// NoContentResponse returns a no content response
func NoContentResponse(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// Helper functions
func getRequestID(c echo.Context) string {
	requestID := c.Response().Header().Get("X-Request-ID")
	if requestID == "" {
		requestID = c.Request().Header.Get("X-Request-ID")
	}
	return requestID
}

func getErrorCode(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "BAD_REQUEST"
	case http.StatusUnauthorized:
		return "UNAUTHORIZED"
	case http.StatusForbidden:
		return "FORBIDDEN"
	case http.StatusNotFound:
		return "NOT_FOUND"
	case http.StatusConflict:
		return "CONFLICT"
	case http.StatusUnprocessableEntity:
		return "VALIDATION_ERROR"
	case http.StatusInternalServerError:
		return "INTERNAL_ERROR"
	case http.StatusServiceUnavailable:
		return "SERVICE_UNAVAILABLE"
	default:
		return "UNKNOWN_ERROR"
	}
}

func formatValidationErrors(errors []ValidationError) string {
	if len(errors) == 0 {
		return "Unknown validation error"
	}

	if len(errors) == 1 {
		return errors[0].Error()
	}

	result := "Multiple validation errors: "
	for i, err := range errors {
		if i > 0 {
			result += "; "
		}
		result += err.Error()
	}
	return result
}
