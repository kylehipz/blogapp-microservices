package main

import "github.com/kylehipz/blogapp-microservices/libs/pkg/api"

func main() {
	apiServer := api.NewEchoAPIServer(":9090")

	routes := []*api.EchoAPIRoute{}
	apiServer.Run(routes)
}
