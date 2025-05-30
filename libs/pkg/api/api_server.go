package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoAPIServer struct {
	e    *echo.Echo
	addr string
}

func NewEchoAPIServer(addr string) *EchoAPIServer {
	e := echo.New()

	return &EchoAPIServer{
		e:    e,
		addr: addr,
	}
}

// Starts the API Server
func (a *EchoAPIServer) Run(prefixPath string, routes []*EchoAPIRoute) {
	a.RegisterRoutes(prefixPath, routes)
	a.e.Pre(middleware.RemoveTrailingSlash())
	a.e.Use(middleware.Recover())

	a.e.Logger.Fatal(a.e.Start(a.addr))
}

// Registers the routes
func (a *EchoAPIServer) RegisterRoutes(prefixPath string, routes []*EchoAPIRoute) {
	grp := a.e.Group(prefixPath)

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			grp.GET(route.Path, route.Handler, route.Middlewares...)
		case http.MethodPost:
			grp.POST(route.Path, route.Handler, route.Middlewares...)
		case http.MethodPatch:
			grp.PATCH(route.Path, route.Handler, route.Middlewares...)
		case http.MethodDelete:
			grp.DELETE(route.Path, route.Handler, route.Middlewares...)
		}
	}
}

func (a *EchoAPIServer) Use(middlewares []echo.MiddlewareFunc) {
	for _, middleware := range middlewares {
		a.e.Use(middleware)
	}
}
