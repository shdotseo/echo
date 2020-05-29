package echo

import (
	"fmt"
	"log"
)

type Echo struct {
	Logger log.Logger
}

func New() (e *Echo) {
	e = &Echo{}
	return
}

func (e *Echo) Start(addr string) error {
	fmt.Printf("%s\n", addr)
	return nil
}

func (e *Echo) GET(url string, arg ...interface{}) error {
	fmt.Printf("GET %s\n", url)
	return nil
}
