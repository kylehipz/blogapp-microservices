package main

import (
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
)

func main() {
	e := echo.New()

	routes := []*api.EchoAPIRoute{}

	apiServer := api.NewEchoAPIServer(":9090")
	apiServer.Run(e, routes)
}
