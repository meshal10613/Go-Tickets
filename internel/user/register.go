package user

import (
	"go-tickets/internel/validation"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	validate := validator.New()
	validate.RegisterValidation("password", validation.PasswordValidation)

	e.Validator = validation.NewCustomValidator(validate)

	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	api := e.Group("/api/v1")
	api.POST("/users", userHandler.CreateUser)
}