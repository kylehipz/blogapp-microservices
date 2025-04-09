package main

import (
	"fmt"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"

	"github.com/kylehipz/blogapp-microservices/blogs/internal"
)

func main() {
	apiServerPort := fmt.Sprintf(":%s", internal.API_SERVER_PORT)
	apiServer := api.NewEchoAPIServer(apiServerPort)

	routes := []*api.EchoAPIRoute{}
	apiServer.Run(routes)
}
