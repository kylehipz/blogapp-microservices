package api

import "github.com/labstack/echo/v4"

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
func (a *EchoAPIServer) Run(routes []*EchoAPIRoute) {
	a.RegisterRoutes(routes)
	a.e.Logger.Fatal(a.e.Start(a.addr))
}

// Registers the routes
func (a *EchoAPIServer) RegisterRoutes(routes []*EchoAPIRoute) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			a.e.GET(route.Path, route.Handler, route.Middlewares...)
		case "POST":
			a.e.POST(route.Path, route.Handler, route.Middlewares...)
		case "PATCH":
			a.e.PATCH(route.Path, route.Handler, route.Middlewares...)
		case "DELETE":
			a.e.DELETE(route.Path, route.Handler, route.Middlewares...)
		}
	}
}
