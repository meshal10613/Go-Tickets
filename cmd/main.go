package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

// func (cv *CustomValidator) Validate(i any) error {
// 	return cv.validator.Struct(i)
// }

var (
	upperRegex   = regexp.MustCompile(`[A-Z]`)
	lowerRegex   = regexp.MustCompile(`[a-z]`)
	numberRegex  = regexp.MustCompile(`[0-9]`)
	specialRegex = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?]`)
)

func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return len(password) >= 8 &&
		upperRegex.MatchString(password) &&
		lowerRegex.MatchString(password) &&
		numberRegex.MatchString(password) &&
		specialRegex.MatchString(password)
}

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name" query:"name" validate:"required" gorm:"type:varchar(100) not null"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email" gorm:"type:varchar(100); uniqueIndex; not null"`
	// Password string `json:"password" form:"password" query:"password" validate:"required,min=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789,containsany=!@#$%^&*()_+-=[]{}|;:,.<>?/"`
	Password string `json:"password" form:"password" query:"password" validate:"required,password" gorm:"type:varchar(255) not null"`
}

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
	db.AutoMigrate(&User{})

	port := 5000
	e := echo.New()
	e.Use(middleware.RequestLogger())
	validate := validator.New()
	validate.RegisterValidation("password", passwordValidation)

	e.Validator = &CustomValidator{
		validator: validate,
	}

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", func(c *echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})
		}

		if err := c.Validate(u); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				errors := make(map[string]string)

				for _, fieldErr := range validationErrors {
					errors[fieldErr.Field()] = fieldErr.Tag()
				}

				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "Validation failed",
					"errors":  errors,
				})
			}

			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
		}

		//? save to database
		result := db.Create(&u)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": result.Error.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "User created successfully",
			"result": map[string]any{
				"id":    u.ID,
				"name":  u.Name,
				"email": u.Email,
			},
		})
	})

	red := color.New(color.FgRed).SprintFunc()

	color.Green("🚀 Server is starting on port http://localhost:%d \n", port)

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		color.Red("❌ Failed to start server: %v", err)
		e.Logger.Error(red(err.Error()))
	}
}
