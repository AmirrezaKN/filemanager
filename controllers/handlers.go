package controllers

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func SendFile(app echo.Context) error {
	id := app.Param("id")
	filepath := fmt.Sprintf("%s/temp/%s", os.Getenv("PWD"), id)
	return app.Attachment(filepath, id)
}
