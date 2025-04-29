package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/cache"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/home-feed/internal/services"
)

func New(conn *pgx.Conn, rdb *redis.Client) []*api.EchoAPIRoute {
	postgresClient := db.NewPostgresClient(conn)
	redisClient := cache.NewRedisClient(rdb)
	homeFeedService := services.NewHomeFeedService(postgresClient, redisClient)

	homeFeedHandler := handlers.NewHomeFeedHandler(homeFeedService)

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
