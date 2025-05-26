package routes

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/cache"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/home-feed/internal/services"
)

func New(
	conn *pgx.Conn,
	rdb *redis.Client,
	rabbitMQClient *pubsub.RabbitMQClient,
	logger *zap.Logger,
) []*api.EchoAPIRoute {
	postgresClient := db.NewPostgresClient(conn)
	redisClient := cache.NewRedisClient(rdb)

	homeFeedService := services.NewHomeFeedService(postgresClient, redisClient, rabbitMQClient)

	homeFeedHandler := handlers.NewHomeFeedHandler(homeFeedService, logger)
	homeFeedEventsHandler := handlers.NewHomeFeedEventsHandler(homeFeedService)

	go func() {
		fmt.Println("Home feed service listening to events")

		homeFeedEventsHandler.StartListener()
	}()

	routes := []*api.EchoAPIRoute{
		{
			Path:        "",
			Method:      http.MethodGet,
			Handler:     homeFeedHandler.GetHomeFeed,
			Middlewares: []echo.MiddlewareFunc{},
		},
	}

	return routes
}
