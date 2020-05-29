package middleware

import (
	"fmt"
	"github.com/shdotseo/echo"
)

func Logger() echo.MiddlewareFunc {
	fmt.Printf("Logger()\n")
	return nil
}
