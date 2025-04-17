package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/follow/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/follow/internal/services"
)

func New(conn *pgx.Conn) []*api.EchoAPIRoute {
	queries := db.New(conn)
	followService := services.FollowService{Queries: queries}
	followHandler := handlers.FollowHandler{FollowService: followService}

	routes := []*api.EchoAPIRoute{
		{
			Path:        "/login",
			Method:      http.MethodPost,
			Handler:     followHandler.FollowUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/register",
			Method:      http.MethodPost,
			Handler:     followHandler.UnfollowUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
	}

	return routes
}
