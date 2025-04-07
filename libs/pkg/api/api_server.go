package api

import "github.com/labstack/echo/v4"

type EchoAPIServer struct {
	addr string
}

type EchoAPIRoute struct {
	method      string
	path        string
	handler     echo.HandlerFunc
	middlewares []echo.MiddlewareFunc
}

func NewEchoAPIServer(addr string) *EchoAPIServer {
	return &EchoAPIServer{
		addr: addr,
	}
}

// Starts the API Server
func (a *EchoAPIServer) Run(e *echo.Echo, routes []*EchoAPIRoute) {
	a.RegisterRoutes(e, routes)
	e.Logger.Fatal(e.Start(a.addr))
}

// Registers the routes
func (a *EchoAPIServer) RegisterRoutes(e *echo.Echo, routes []*EchoAPIRoute) {
	for _, route := range routes {
		switch route.method {
		case "GET":
			e.GET(route.path, route.handler, route.middlewares...)
		case "POST":
			e.POST(route.path, route.handler, route.middlewares...)
		case "PATCH":
			e.PATCH(route.path, route.handler, route.middlewares...)
		case "DELETE":
			e.DELETE(route.path, route.handler, route.middlewares...)
		}
	}
}
