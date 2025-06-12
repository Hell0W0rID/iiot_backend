package cors

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS returns CORS middleware with appropriate configuration for IoT backend
func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:5000",
			"http://localhost:8080",
			"https://localhost:3000",
			"https://localhost:3001",
			"https://localhost:5000",
			"https://localhost:8080",
		},
		AllowMethods: []string{
			echo.GET,
			echo.HEAD,
			echo.PUT,
			echo.PATCH,
			echo.POST,
			echo.DELETE,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-API-Key",
			"X-Requested-With",
			"X-CSRF-Token",
			"X-Request-ID",
		},
		ExposeHeaders: []string{
			"X-Request-ID",
			"X-Response-Time",
			"X-Total-Count",
		},
		AllowCredentials: true,
		MaxAge:          300, // 5 minutes
	})
}

// RestrictiveCORS returns a more restrictive CORS configuration for production
func RestrictiveCORS(allowedOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-API-Key",
		},
		ExposeHeaders: []string{
			"X-Request-ID",
		},
		AllowCredentials: false,
		MaxAge:          86400, // 24 hours
	})
}

// DevCORS returns a permissive CORS configuration for development
func DevCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:          300,
	})
}
