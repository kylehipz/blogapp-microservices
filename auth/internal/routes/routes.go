package routes

import (
	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
)

func New(conn *pgx.Conn) []*api.EchoAPIRoute {
	queries := db.New(conn)
	usersService := services.UsersService{Queries: queries}
	usersHandler := handlers.UsersHandler{UsersService: usersService}

	routes := []*api.EchoAPIRoute{
		{
			Path:        "/login",
			Method:      "GET",
			Handler:     usersHandler.LoginUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/register",
			Method:      "GET",
			Handler:     usersHandler.RegisterUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
	}

	return routes
}
