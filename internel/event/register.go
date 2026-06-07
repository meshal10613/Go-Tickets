package event

import (
	"go-tickets/internel/auth"
	"go-tickets/internel/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)
	jwtService := auth.NewJwtService("")

	api := e.Group("/api/v1/event")
	api.POST("", handler.Create, middlewares.AuthMiddleware(jwtService))
}
