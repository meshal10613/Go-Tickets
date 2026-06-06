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

	userRepository := NewUserRepository(db)
	jwtService := auth.NewJwtService("") //? You can pass secret key from config
	userService := NewUserService(userRepository, jwtService)
	userHandler := NewUserHandler(userService)

	api := e.Group("/api/v1/auth")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/me", userHandler.GetMe, middlewares.AuthMiddleware(jwtService))
}
