package main

import (
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
)

func main() {
	routes := []*api.EchoAPIRoute{}

	apiServer := api.NewEchoAPIServer(":9090")
	apiServer.Run("/libs", routes)
}
