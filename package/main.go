package main

// ทดสอบการเปิด Server echo framwork

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go !")
	})
	e.Logger.Fatal(e.Start(":1323"))
}