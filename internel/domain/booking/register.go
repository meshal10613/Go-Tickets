package booking

import (
	"go-tickets/internel/auth"
	"go-tickets/internel/domain/event"
	"go-tickets/internel/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	bookingRepo := NewRepository(db)
	eventRepo := event.NewRepository(db)

	service := NewService(bookingRepo, eventRepo)
	handler := NewHandler(service)

	jwtService := auth.NewJwtService("")

	api := e.Group("/api/v1/booking", middlewares.AuthMiddleware(jwtService))

	api.POST("", handler.Create)
	api.GET("/my-booking", handler.GetMyBookings)
}
