package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Start(":8000")
	fmt.Println("success")
}
