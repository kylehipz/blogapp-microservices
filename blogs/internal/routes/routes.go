package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/api"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/blogs/internal/handlers"
	"github.com/kylehipz/blogapp-microservices/blogs/internal/services"
)

func New(conn *pgx.Conn, rabbitMQClient *pubsub.RabbitMQClient) []*api.EchoAPIRoute {
	postgresClient := db.NewPostgresClient(conn)
	blogsService := services.NewBlogsService(postgresClient, rabbitMQClient)
	blogsHandler := handlers.NewBlogsHandler(blogsService)

	routes := []*api.EchoAPIRoute{
		{
			Path:        "",
			Method:      http.MethodPost,
			Handler:     blogsHandler.CreateBlog,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/:id",
			Method:      http.MethodGet,
			Handler:     blogsHandler.GetBlog,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/:id",
			Method:      http.MethodPatch,
			Handler:     blogsHandler.UpdateBlog,
			Middlewares: []echo.MiddlewareFunc{},
		},
		{
			Path:        "/:id",
			Method:      http.MethodDelete,
			Handler:     blogsHandler.DeleteBlog,
			Middlewares: []echo.MiddlewareFunc{},
		},
	}

	return routes
}
