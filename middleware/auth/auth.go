package auth

import (
        "net/http"
        "os"

        "github.com/golang-jwt/jwt/v5"
        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
)

// JWTAuth returns JWT authentication middleware
func JWTAuth() echo.MiddlewareFunc {
        // Get JWT secret from environment
        secret := os.Getenv("JWT_SECRET")
        if secret == "" {
                secret = "iiot-backend-secret-key" // fallback for development
        }

        // Skip authentication for health check and certain public endpoints
        skipper := func(c echo.Context) bool {
                path := c.Request().URL.Path
                method := c.Request().Method

                // Skip authentication for these paths
                skipPaths := []string{
                        "/health",
                        "/api/v2/ping",
                }

                // Skip for OPTIONS requests (CORS preflight)
                if method == "OPTIONS" {
                        return true
                }

                for _, skipPath := range skipPaths {
                        if path == skipPath {
                                return true
                        }
                }

                // Check if API key authentication is enabled and provided
                apiKey := c.Request().Header.Get("X-API-Key")
                if apiKey != "" {
                        expectedAPIKey := os.Getenv("API_KEY")
                        if expectedAPIKey != "" && apiKey == expectedAPIKey {
                                return true // Skip JWT if valid API key is provided
                        }
                }

                return false
        }

        config := middleware.JWTConfig{
                SigningKey:  []byte(secret),
                TokenLookup: "header:Authorization:Bearer ,query:token,cookie:token",
                Skipper:     skipper,
        }

        return middleware.JWTWithConfig(config)
}

// APIKeyAuth returns API key authentication middleware
func APIKeyAuth() echo.MiddlewareFunc {
        return func(next echo.HandlerFunc) echo.HandlerFunc {
                return func(c echo.Context) error {
                        // Skip for health check and OPTIONS requests
                        if c.Request().URL.Path == "/health" || c.Request().Method == "OPTIONS" {
                                return next(c)
                        }

                        // Check for API key in header
                        apiKey := c.Request().Header.Get("X-API-Key")
                        if apiKey == "" {
                                // Also check in query parameter
                                apiKey = c.QueryParam("api_key")
                        }

                        if apiKey == "" {
                                return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
                                        "error":   "Unauthorized",
                                        "message": "API key is required",
                                        "code":    "API_KEY_REQUIRED",
                                })
                        }

                        // Validate API key
                        expectedAPIKey := os.Getenv("API_KEY")
                        if expectedAPIKey == "" {
                                // If no API key is configured, allow the request
                                return next(c)
                        }

                        if apiKey != expectedAPIKey {
                                return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
                                        "error":   "Unauthorized",
                                        "message": "Invalid API key",
                                        "code":    "INVALID_API_KEY",
                                })
                        }

                        return next(c)
                }
        }
}

// BasicAuth returns basic authentication middleware
func BasicAuth(username, password string) echo.MiddlewareFunc {
        return middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
                return u == username && p == password, nil
        })
}

// ExtractUserFromToken extracts user information from JWT token
func ExtractUserFromToken(c echo.Context) (map[string]interface{}, error) {
        token, ok := c.Get("user").(*jwt.Token)
        if !ok {
                return nil, echo.NewHTTPError(http.StatusUnauthorized, "No valid token found")
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
                return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
        }

        user := make(map[string]interface{})
        for key, value := range claims {
                user[key] = value
        }

        return user, nil
}

// RequireRole middleware to check for specific roles
func RequireRole(requiredRole string) echo.MiddlewareFunc {
        return func(next echo.HandlerFunc) echo.HandlerFunc {
                return func(c echo.Context) error {
                        user, err := ExtractUserFromToken(c)
                        if err != nil {
                                return err
                        }

                        role, ok := user["role"].(string)
                        if !ok || role != requiredRole {
                                return echo.NewHTTPError(http.StatusForbidden, map[string]interface{}{
                                        "error":   "Forbidden",
                                        "message": "Insufficient permissions",
                                        "code":    "INSUFFICIENT_PERMISSIONS",
                                })
                        }

                        return next(c)
                }
        }
}

// RequireAnyRole middleware to check for any of the specified roles
func RequireAnyRole(requiredRoles ...string) echo.MiddlewareFunc {
        return func(next echo.HandlerFunc) echo.HandlerFunc {
                return func(c echo.Context) error {
                        user, err := ExtractUserFromToken(c)
                        if err != nil {
                                return err
                        }

                        role, ok := user["role"].(string)
                        if !ok {
                                return echo.NewHTTPError(http.StatusForbidden, map[string]interface{}{
                                        "error":   "Forbidden",
                                        "message": "No role found in token",
                                        "code":    "NO_ROLE_FOUND",
                                })
                        }

                        for _, requiredRole := range requiredRoles {
                                if role == requiredRole {
                                        return next(c)
                                }
                        }

                        return echo.NewHTTPError(http.StatusForbidden, map[string]interface{}{
                                "error":   "Forbidden",
                                "message": "Insufficient permissions",
                                "code":    "INSUFFICIENT_PERMISSIONS",
                                "required_roles": requiredRoles,
                                "user_role": role,
                        })
                }
        }
}
