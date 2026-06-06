package middlewares

import (
	"go-tickets/internel/auth"
	httpsresponse "go-tickets/internel/httpsResponse"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JwtService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			//? extract token from authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpsresponse.Error{
					Success: false,
					Message: "Missing authorization header!",
				})
			}

			//? check bearer scheme
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, httpsresponse.Error{
					Success: false,
					Message: "Invalid authorization header!",
				})
			}
			tokenString := parts[1]

			//? validate token
			claims, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpsresponse.Error{
					Success: false,
					Message: "Invalid or expired token!",
				})
			}

			//? store user info in context for handler
			c.Set("user_id", claims.UserID)
			c.Set("name", claims.Name)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}
