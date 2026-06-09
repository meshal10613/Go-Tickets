package middlewares

import (
	"go-tickets/internel/auth"
	httpsresponse "go-tickets/internel/httpResponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JwtService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx *echo.Context) error {
			// //? extract token from authorization header
			// authHeader := ctx.Request().Header.Get("Authorization")
			// if authHeader == "" {
			// 	return ctx.JSON(http.StatusUnauthorized, httpsresponse.Error{
			// 		Success: false,
			// 		Message: "Authorization header is required",
			// 	})
			// }

			// //? check bearer scheme
			// parts := strings.Split(authHeader, " ")
			// if len(parts) != 2 || parts[0] != "Bearer" {
			// 	return ctx.JSON(http.StatusUnauthorized, httpsresponse.Error{
			// 		Success: false,
			// 		Message: "Invalid authorization header format. Expected: Bearer <token>",
			// 	})
			// }
			// tokenString := parts[1]

			cookie, err := ctx.Cookie("token")
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, httpsresponse.Error{
					Success: false,
					Message: "Authentication required",
				})
			}

			tokenString := cookie.Value

			//? validate token
			claims, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, httpsresponse.Error{
					Success: false,
					Message: "Authentication failed. Invalid or expired token",
				})
			}

			//? store user info in context for handler
			ctx.Set("user_id", claims.UserID)
			ctx.Set("name", claims.Name)
			ctx.Set("email", claims.Email)

			return next(ctx)
		}
	}
}
