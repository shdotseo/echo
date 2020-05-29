package main

import (
	"github.com/shdotseo/echo"
	"github.com/shdotseo/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
