package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/loadenv"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/kylehipz/blogapp-microservices/follow/internal/routes"
)

func main() {
	// load .env
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" || environment == "development" {
		loadenv.Load()
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// start database
	ctx := context.Background()

	// start database
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("database connection error", zap.Error(err))
	}
	defer conn.Close(ctx)

	logger.Info("database connected")

	// start api server
	apiServerPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	apiServer := api.NewEchoAPIServer(apiServerPort)

	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(os.Getenv("JWT_SECRET"))
	apiServer.Use([]echo.MiddlewareFunc{authenticationMiddleware})

	followRoutes := routes.New(conn)
	apiServer.Run("/follow", followRoutes)
}
