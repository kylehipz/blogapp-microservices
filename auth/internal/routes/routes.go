package routes

import (
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
)

func New(conn *pgx.Conn) []*api.EchoAPIRoute {
	queries := db.New(conn)
	usersService := services.UsersService{Queries: queries}
	usersHandler := handlers.UsersHandler{UsersService: usersService}

	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(os.Getenv("JWT_SECRET"))

	routes := []*api.EchoAPIRoute{
		{
			Path:        "/login",
			Method:      http.MethodPost,
			Handler:     usersHandler.LoginUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/register",
			Method:      http.MethodPost,
			Handler:     usersHandler.RegisterUser,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/test",
			Method:      http.MethodGet,
			Handler:     usersHandler.TestRoute,
			Middlewares: []echo.MiddlewareFunc{authenticationMiddleware},
		},
	}

	return routes
}
