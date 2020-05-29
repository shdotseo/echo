package echo

import "fmt"

type Context struct{}

func (c *Context) String(status int, message string) error {
	fmt.Printf("%d %s\n", status, message)
	return nil
}
