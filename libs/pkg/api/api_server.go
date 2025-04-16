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
func (a *EchoAPIServer) Run(prefixPath string, routes []*EchoAPIRoute) {
	a.RegisterRoutes(prefixPath, routes)
	a.e.Logger.Fatal(a.e.Start(a.addr))
}

// Registers the routes
func (a *EchoAPIServer) RegisterRoutes(prefixPath string, routes []*EchoAPIRoute) {
	grp := a.e.Group(prefixPath)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			grp.GET(route.Path, route.Handler, route.Middlewares...)
		case "POST":
			grp.POST(route.Path, route.Handler, route.Middlewares...)
		case "PATCH":
			grp.PATCH(route.Path, route.Handler, route.Middlewares...)
		case "DELETE":
			grp.DELETE(route.Path, route.Handler, route.Middlewares...)
		}
	}
}
