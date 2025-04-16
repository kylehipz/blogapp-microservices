package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/loadenv"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
)

func main() {
	// load .env
	loadenv.Load()

	// start database
	ctx := context.Background()

	// start database
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	log.Println("Successfully connected to the database")

	// initiate handlers and services
	queries := db.New(conn)
	usersService := services.UsersService{Queries: queries}
	usersHandler := handlers.UsersHandler{UsersService: usersService}

	// start API server
	apiServerPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	apiServer := api.NewEchoAPIServer(apiServerPort)

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
	apiServer.Run("/auth", routes)
}
