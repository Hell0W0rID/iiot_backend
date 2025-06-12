// Package main provides the EdgeX Core Data service main entry point
// following edgexfoundry/edgex-go architecture patterns
package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"iiot-backend/services/core/data/application"
	"iiot-backend/pkg/common"
)

func main() {
	// Load database configuration from environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:password@localhost/postgres?sslmode=disable"
	}

	// Initialize database connection
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		panic("Database connection failed: " + err.Error())
	}

	// Create Echo instance for EdgeX Core Data
	e := echo.New()
	e.HideBanner = true

	// Add EdgeX standard middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize EdgeX Core Data service
	dataService := application.NewDataService()

	// EdgeX v3 API routes for Core Data
	v3 := e.Group("/api/v3")

	// Event routes (following EdgeX patterns)
	v3.GET("/event/all", func(c echo.Context) error {
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		deviceName := c.QueryParam("device")
		
		response, edgeErr := dataService.AllEvents(c.Request().Context(), offset, limit, deviceName)
		if edgeErr.Code() != 0 {
			return c.JSON(edgeErr.Code(), map[string]string{"error": edgeErr.Message()})
		}
		
		return c.JSON(http.StatusOK, response)
	})

	v3.GET("/event/id/:id", func(c echo.Context) error {
		id := c.Param("id")
		
		response, edgeErr := dataService.EventById(c.Request().Context(), id)
		if edgeErr.Code() != 0 {
			return c.JSON(edgeErr.Code(), map[string]string{"error": edgeErr.Message()})
		}
		
		return c.JSON(http.StatusOK, response)
	})

	v3.DELETE("/event/id/:id", func(c echo.Context) error {
		id := c.Param("id")
		
		response, edgeErr := dataService.DeleteEventById(c.Request().Context(), id)
		if edgeErr.Code() != 0 {
			return c.JSON(edgeErr.Code(), map[string]string{"error": edgeErr.Message()})
		}
		
		return c.JSON(http.StatusOK, response)
	})

	// Reading routes
	v3.GET("/reading/device/name/:name", func(c echo.Context) error {
		deviceName := c.Param("name")
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		
		response, edgeErr := dataService.ReadingsByDeviceName(c.Request().Context(), deviceName, offset, limit)
		if edgeErr.Code() != 0 {
			return c.JSON(edgeErr.Code(), map[string]string{"error": edgeErr.Message()})
		}
		
		return c.JSON(http.StatusOK, response)
	})

	v3.GET("/reading/resource/name/:name", func(c echo.Context) error {
		resourceName := c.Param("name")
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		
		response, edgeErr := dataService.ReadingsByResourceName(c.Request().Context(), resourceName, offset, limit)
		if edgeErr.Code() != 0 {
			return c.JSON(edgeErr.Code(), map[string]string{"error": edgeErr.Message()})
		}
		
		return c.JSON(http.StatusOK, response)
	})

	// EdgeX common routes
	e.GET("/api/v3/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"apiVersion":  common.ApiVersion,
			"timestamp":   common.GetTimestamp(),
			"serviceName": common.CoreDataServiceKey,
		})
	})

	e.GET("/api/v3/version", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"apiVersion":  common.ApiVersion,
			"version":     common.ServiceVersion,
			"serviceName": common.CoreDataServiceKey,
			"sdkVersion":  "3.0.0",
		})
	})

	// EdgeX configuration endpoint
	e.GET("/api/v3/config", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"apiVersion":  common.ApiVersion,
			"serviceName": common.CoreDataServiceKey,
			"config": map[string]interface{}{
				"service": map[string]interface{}{
					"host":            "0.0.0.0",
					"port":            59880,
					"startupMsg":      "EdgeX Core Data service has started",
					"maxResultCount":  1000,
					"requestTimeout":  "20s",
				},
				"database": map[string]interface{}{
					"type": "postgres",
					"host": "localhost",
					"port": 5432,
				},
				"retention": map[string]interface{}{
					"enabled":       true,
					"interval":      "30s",
					"maxCap":        10000,
					"minTimeToLive": "1h",
				},
			},
		})
	})

	// Start server on EdgeX Core Data standard port
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "59880" // EdgeX Core Data standard port
	}

	// Graceful shutdown handling
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := e.Start(":" + port); err != nil {
			cancel()
		}
	}()

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-c:
		cancel()
	case <-ctx.Done():
	}

	// Shutdown server
	e.Shutdown(ctx)
	wg.Wait()
}