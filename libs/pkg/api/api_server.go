package api

import "github.com/labstack/echo/v4"

type EchoAPIServer struct {
	e    *echo.Echo
	addr string
}

type EchoAPIRoute struct {
	method      string
	path        string
	handler     echo.HandlerFunc
	middlewares []echo.MiddlewareFunc
}

func NewEchoAPIServer(addr string) *EchoAPIServer {
	e := echo.New()

	return &EchoAPIServer{
		e:    e,
		addr: addr,
	}
}

// Starts the API Server
func (a *EchoAPIServer) Run(routes []*EchoAPIRoute) {
	a.RegisterRoutes(routes)
	a.e.Logger.Fatal(a.e.Start(a.addr))
}

// Registers the routes
func (a *EchoAPIServer) RegisterRoutes(routes []*EchoAPIRoute) {
	for _, route := range routes {
		switch route.method {
		case "GET":
			a.e.GET(route.path, route.handler, route.middlewares...)
		case "POST":
			a.e.POST(route.path, route.handler, route.middlewares...)
		case "PATCH":
			a.e.PATCH(route.path, route.handler, route.middlewares...)
		case "DELETE":
			a.e.DELETE(route.path, route.handler, route.middlewares...)
		}
	}
}
