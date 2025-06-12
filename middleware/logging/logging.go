package logging

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Logger returns logging middleware with custom configuration
func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			// Skip logging for health checks to reduce noise
			return c.Request().URL.Path == "/health"
		},
		Format: `{"time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}",` +
			`"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	})
}

// RequestLogger logs detailed request information
func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Process request
			err := next(c)

			// Calculate latency
			latency := time.Since(start)

			// Get request and response information
			req := c.Request()
			res := c.Response()

			// Log the request details
			log.Infof("Request processed - Method: %s, Path: %s, Status: %d, Latency: %v, RemoteIP: %s, UserAgent: %s",
				req.Method,
				req.URL.Path,
				res.Status,
				latency,
				c.RealIP(),
				req.UserAgent(),
			)

			// Log errors with more detail
			if err != nil {
				log.Errorf("Request error - Method: %s, Path: %s, Error: %v, RemoteIP: %s",
					req.Method,
					req.URL.Path,
					err,
					c.RealIP(),
				)
			}

			return err
		}
	}
}

// ErrorLogger logs errors with detailed context
func ErrorLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				// Log error with context
				log.Errorf("Error occurred - Method: %s, Path: %s, Error: %v, RemoteIP: %s, UserAgent: %s",
					c.Request().Method,
					c.Request().URL.Path,
					err,
					c.RealIP(),
					c.Request().UserAgent(),
				)
			}
			return err
		}
	}
}

// StructuredLogger returns a structured logging middleware
func StructuredLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			req := c.Request()
			res := c.Response()
			latency := time.Since(start)

			// Create structured log entry
			logEntry := map[string]interface{}{
				"timestamp":   start.Format(time.RFC3339),
				"method":      req.Method,
				"path":        req.URL.Path,
				"query":       req.URL.RawQuery,
				"status":      res.Status,
				"latency_ms":  latency.Milliseconds(),
				"remote_ip":   c.RealIP(),
				"user_agent":  req.UserAgent(),
				"referer":     req.Referer(),
				"bytes_in":    req.ContentLength,
				"bytes_out":   res.Size,
			}

			// Add error information if present
			if err != nil {
				logEntry["error"] = err.Error()
				log.Errorj(logEntry)
			} else {
				// Log successful requests at debug level for health checks
				if req.URL.Path == "/health" {
					log.Debugj(logEntry)
				} else {
					log.Infoj(logEntry)
				}
			}

			return err
		}
	}
}

// SetLogLevel sets the global log level
func SetLogLevel(level string) {
	switch level {
	case "debug":
		log.SetLevel(log.DEBUG)
	case "info":
		log.SetLevel(log.INFO)
	case "warn":
		log.SetLevel(log.WARN)
	case "error":
		log.SetLevel(log.ERROR)
	default:
		log.SetLevel(log.INFO)
	}
}
