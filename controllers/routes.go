package controllers

import "github.com/labstack/echo/v4"

// SetupRoutes sets up the routes
func SetupRoutes(app *echo.Echo) {
	(*app).GET("/", nil)
	(*app).GET("/file/:id", SendFile)
}
