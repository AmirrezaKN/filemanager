package controllers

import "github.com/labstack/echo/v4"

func setupRoutes(app *echo.Echo) {
	(*app).GET("/", nil)
}
