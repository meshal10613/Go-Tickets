package main

import (
	"fmt"
	"go-tickets/internel/user"
	"go-tickets/internel/validation"

	// "go-tickets/internal/validator"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=4219 dbname=go-tickets port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		// panic("failed to connect database")
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	fmt.Println("Database connected successfully")
	db.AutoMigrate(&user.User{})

	port := 5000
	e := echo.New()
	e.Use(middleware.RequestLogger())
	validate := validator.New()
	validate.RegisterValidation("password", validation.PasswordValidation)

	e.Validator = validation.NewCustomValidator(validate)

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	e.POST("/users", userHandler.CreateUser)

	red := color.New(color.FgRed).SprintFunc()

	color.Green("🚀 Server is starting on port http://localhost:%d \n", port)

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		color.Red("❌ Failed to start server: %v", err)
		e.Logger.Error(red(err.Error()))
	}
}
