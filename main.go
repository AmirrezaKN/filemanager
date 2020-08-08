package main

import (
	"app/controllers"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	controllers.SetupRoutes(app)
	app.Start(":8000")
	fmt.Println("success")
}
