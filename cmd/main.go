package main

import (
	//"github.com/labstack/echo/v4"
	"github.com/shdotseo/echo"
	"net/http"
	//"net/http"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	//e.Logger.Fatal(e.Start(":1323"))
	e.Start(":1323")
}
