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
	"github.com/redis/go-redis/v9"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/routes"
)

func main() {
	// load .env
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" || environment == "development" {
		loadenv.Load()
	}

	ctx := context.Background()

	// connect to postgres database
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	log.Println("Successfully connected to the database")

	// connect to redis
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_URL")})

	apiServerPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	apiServer := api.NewEchoAPIServer(apiServerPort)

	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(os.Getenv("JWT_SECRET"))

	apiServer.Use([]echo.MiddlewareFunc{authenticationMiddleware})

	// connect to rabbitmq
	rabbitConn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		panic(err)
	}

	ch, err := rabbitConn.Channel()
	if err != nil {
		panic(err)
	}

	rabbitMQClient := pubsub.NewRabbitMQClient(rabbitConn, ch, "blogapp", "home-feed", "home-feed")

	defer rabbitMQClient.CleanUp()

	homeFeedRoutes := routes.New(conn, rdb, rabbitMQClient)

	apiServer.Run("/home-feed", homeFeedRoutes)
}
