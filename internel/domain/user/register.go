package user

import (
	"go-tickets/internel/auth"
	"go-tickets/internel/middlewares"
	"go-tickets/internel/validation"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	validate := validator.New()
	validate.RegisterValidation("password", validation.PasswordValidation)

	e.Validator = validation.NewCustomValidator(validate)

	repository := NewRepository(db)
	jwtService := auth.NewJwtService("") //? You can pass secret key from config
	service := NewService(repository, jwtService)
	handler := NewHandler(service)

	api := e.Group("/api/v1/auth")
	api.POST("/register", handler.RegisterUser)
	api.POST("/login", handler.LoginUser)
	api.GET("/me", handler.GetMe, middlewares.AuthMiddleware(jwtService))
}
