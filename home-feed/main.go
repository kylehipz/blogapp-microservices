package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal"
)

func main() {
	ctx := context.Background()

	// start database
	conn, err := pgx.Connect(ctx, internal.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	log.Println("Successfully connected to the database")

	db.New(conn)

	apiServerPort := fmt.Sprintf(":%s", internal.API_SERVER_PORT)
	apiServer := api.NewEchoAPIServer(apiServerPort)

	routes := []*api.EchoAPIRoute{}
	apiServer.Run("/home-feed", routes)
}
