package routes

import (
	"github.com/labstack/echo/v4"
)


func Router() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":4320"))
}