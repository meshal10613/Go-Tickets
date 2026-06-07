package server

import (
	"go-tickets/internel/config"
	"go-tickets/internel/event"
	httpresponse "go-tickets/internel/httpResponse"
	"go-tickets/internel/user"
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB, cfg *config.Config) {
	e := echo.New()
	db.AutoMigrate(&user.User{}, &event.Event{})

	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Success{
			Success: true,
			Message: "Server is running successfully....",
		})
	})

	//? User routes
	user.RegisterRoutes(e, db)
	event.RegisterRoutes(e, db)

	red := color.New(color.FgRed).SprintFunc()
	color.Green("🚀 Server running at http://localhost:%s", cfg.Port)

	if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
		color.Red("❌ Failed to start server: %v", err)
		e.Logger.Error(red(err.Error()))
	}
}
