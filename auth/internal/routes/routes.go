package routes

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
	"github.com/kylehipz/blogapp-microservices/auth/internal/types"
)

func New(conn *pgx.Conn) []*api.EchoAPIRoute {
	queries := db.New(conn)
	usersService := services.UsersService{Queries: queries}
	usersHandler := handlers.UsersHandler{UsersService: usersService}

	authConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

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
			Middlewares: []echo.MiddlewareFunc{echojwt.WithConfig(authConfig)},
		},
	}

	return routes
}
