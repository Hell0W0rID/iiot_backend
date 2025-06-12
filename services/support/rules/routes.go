package rules

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, service *Service) {
	handler := NewHandler(service)

	// Rules routes
	rules := g.Group("/rule")
	rules.GET("", handler.GetRules)
	rules.GET("/:id", handler.GetRule)
	rules.GET("/name/:name", handler.GetRuleByName)
	rules.POST("", handler.CreateRule)
	rules.PUT("/:id", handler.UpdateRule)
	rules.DELETE("/:id", handler.DeleteRule)
	rules.PUT("/:id/enable", handler.EnableRule)
	rules.PUT("/:id/disable", handler.DisableRule)
	rules.POST("/:id/execute", handler.ExecuteRule)

	// Pipelines routes
	pipelines := g.Group("/pipeline")
	pipelines.GET("", handler.GetPipelines)
	pipelines.GET("/:id", handler.GetPipeline)
	pipelines.GET("/name/:name", handler.GetPipelineByName)
	pipelines.POST("", handler.CreatePipeline)
	pipelines.PUT("/:id", handler.UpdatePipeline)
	pipelines.DELETE("/:id", handler.DeletePipeline)
	pipelines.PUT("/:id/enable", handler.EnablePipeline)
	pipelines.PUT("/:id/disable", handler.DisablePipeline)

	// Rule execution routes
	g.GET("/rule/execution", handler.GetRuleExecutions)
}
