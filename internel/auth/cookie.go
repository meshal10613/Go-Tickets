package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func SetAuthCookie(ctx *echo.Context, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // false for local development
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 1, // 1 days
	}

	ctx.SetCookie(cookie)
}

func ClearAuthCookie(ctx *echo.Context) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}

	ctx.SetCookie(cookie)
}