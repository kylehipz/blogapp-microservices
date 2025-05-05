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
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/kylehipz/blogapp-microservices/blogs/internal/routes"
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

	rabbitConn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		panic(err)
	}

	ch, err := rabbitConn.Channel()
	if err != nil {
		panic(err)
	}

	rabbitMQClient := pubsub.NewRabbitMQClient(rabbitConn, ch, "blogapp", "blogs", "blogs")

	defer rabbitMQClient.CleanUp()

	blogRoutes := routes.New(conn, rabbitMQClient)

	apiServer.Run("/blogs", blogRoutes)
}
