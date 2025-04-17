package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/loadenv"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/routes"
)

func main() {
	// load .env
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" || environment == "development" {
		loadenv.Load()
	}

	// start database
	ctx := context.Background()

	// start database
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	log.Println("Successfully connected to the database")

	apiServerPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	apiServer := api.NewEchoAPIServer(apiServerPort)

	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(os.Getenv("JWT_SECRET"))

	apiServer.Use([]echo.MiddlewareFunc{authenticationMiddleware})

	homeFeedRoutes := routes.New(conn)

	apiServer.Run("/home-feed", homeFeedRoutes)
}
