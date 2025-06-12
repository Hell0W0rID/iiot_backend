// Package main provides the EdgeX Core Metadata service main entry point
// following edgexfoundry/edgex-go architecture patterns
package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"iiot-backend/services/core/metadata"
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

	// Create Echo instance for EdgeX Core Metadata
	e := echo.New()
	e.HideBanner = true

	// Add EdgeX standard middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize EdgeX Core Metadata service
	metadataService := metadata.NewWorkingMetadataService(db)

	// Register EdgeX v3 API routes
	v3 := e.Group("/api/v3")
	metadata.RegisterWorkingEdgeXRoutes(v3, metadataService)

	// EdgeX common routes
	e.GET("/api/v3/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"apiVersion":  common.ApiVersion,
			"timestamp":   common.GetTimestamp(),
			"serviceName": common.CoreMetadataServiceKey,
		})
	})

	e.GET("/api/v3/version", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"apiVersion":  common.ApiVersion,
			"version":     common.ServiceVersion,
			"serviceName": common.CoreMetadataServiceKey,
			"sdkVersion":  "3.0.0",
		})
	})

	// EdgeX configuration endpoint
	e.GET("/api/v3/config", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"apiVersion":  common.ApiVersion,
			"serviceName": common.CoreMetadataServiceKey,
			"config": map[string]interface{}{
				"service": map[string]interface{}{
					"host":            "0.0.0.0",
					"port":            59881,
					"startupMsg":      "EdgeX Core Metadata service has started",
					"maxResultCount":  1000,
					"requestTimeout":  "20s",
				},
				"database": map[string]interface{}{
					"type": "postgres",
					"host": "localhost",
					"port": 5432,
				},
			},
		})
	})

	// Start server on EdgeX Core Metadata standard port
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "59881" // EdgeX Core Metadata standard port
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