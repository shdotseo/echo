package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
)

type (
	Echo struct {
		Logger    Logger
		colorer   *color.Color
		Server    *http.Server
		TLSServer *http.Server
	}

	HandlerFunc    func(Context) error
	MiddlewareFunc func(HandlerFunc) HandlerFunc

	Route struct {
		Method string `json:"method"`
		Path   string `json:"path"`
		Name   string `json:"name"`
	}
)

func New() (e *Echo) {
	e = &Echo{
		Server:    new(http.Server),
		TLSServer: new(http.Server),
		Logger:    log.New("echo"),
	}
	return
}

func (e *Echo) Use(m MiddlewareFunc) {
	fmt.Printf("Use\n")
}

func (e *Echo) Start(addr string) error {
	e.Server.Addr = addr
	return e.StartServer(e.Server)
}

func (e *Echo) StartServer(s *http.Server) (err error) {
	// Setup
	e.colorer.SetOutput(e.Logger.Output())
	return nil
}

func (e *Echo) Add(method, path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.add("", method, path, h, m)
}

func (e *Echo) GET(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Add(http.MethodGet, path, h, m)
}

func (e *Echo) add(host, method, path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	name := handlerName(handler)
	router := e.findRouter(host)
	router.Add(method, path, func(c Context) error {
		h := handler
		// Chain middleware
		for i := len(middleware) - 1; i >= 0; i-- {
			h = middleware[i](h)
		}
		return h(c)
	})
	r := &Route{
		Method: m,
		Path:   path,
		Name:   name,
	}
	e.router.routes[method+path] = r
	return r
}
