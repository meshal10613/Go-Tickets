package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	red := color.New(color.FgRed).SprintFunc()

	color.Green("🚀 Server is starting on port http://localhost:5000")

	if err := e.Start(":5000"); err != nil {
		color.Red("❌ Failed to start server: %v", err)
		e.Logger.Error(red(err.Error()))
	}
}
