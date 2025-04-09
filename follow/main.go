package main

import (
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"

	"github.com/kylehipz/blogapp-microservices/follow/internal"
)

func main() {
	apiServer := api.NewEchoAPIServer(internal.API_SERVER_PORT)

	routes := []*api.EchoAPIRoute{}
	apiServer.Run(routes)
}
