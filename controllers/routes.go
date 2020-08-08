package controllers

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo) {
	(*app).GET("/", nil)
	(*app).GET("/file/:id", SendFile)
}
