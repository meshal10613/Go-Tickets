package main

import (
	"fmt"
	"go-tickets/internel/config"
	"go-tickets/internel/user"

	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Sprintf("failed to load environment variables: %v", err))
	}
	port := cfg.Port

	db := config.ConnectDatabase(cfg)
	db.AutoMigrate(&user.User{})

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//? User routes
	user.RegisterRoutes(e, db)

	red := color.New(color.FgRed).SprintFunc()

	color.Green("🚀 Server is starting on port http://localhost:%s \n", port)

	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		color.Red("❌ Failed to start server: %v", err)
		e.Logger.Error(red(err.Error()))
	}
}
