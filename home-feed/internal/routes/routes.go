package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/home-feed/internal/services"
)

func New(conn *pgx.Conn) []*api.EchoAPIRoute {
	queries := db.New(conn)
	homeFeedService := services.HomeFeedService{Queries: queries}
	homeFeedHandler := handlers.HomeFeedHandler{HomeFeedService: &homeFeedService}

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
