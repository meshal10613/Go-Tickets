package server

import (
	"fmt"
	"go-tickets/internel/config"
	"go-tickets/internel/user"
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB, cfg *config.Config) {
	e := echo.New()
	db.AutoMigrate(&user.User{})

	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//? User routes
	user.RegisterRoutes(e, db)

	red := color.New(color.FgRed).SprintFunc()

	color.Green("🚀 Server is starting on port http://localhost:%s \n", cfg.Port)

	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		color.Red("❌ Failed to start server: %v", err)
		e.Logger.Error(red(err.Error()))
	}
}
