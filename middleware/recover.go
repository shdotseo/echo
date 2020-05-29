package middleware

import (
	"fmt"
	"github.com/shdotseo/echo"
)

func Recover() echo.MiddlewareFunc {
	fmt.Printf("Recover()\n")
	return nil
}
