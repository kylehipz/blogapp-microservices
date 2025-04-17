package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/blogs/internal/services"
)

type BlogsHandler struct {
	BlogsService *services.BlogsService
}

func (b *BlogsHandler) CreateBlog(c echo.Context) error {
	return c.JSON(http.StatusCreated, echo.Map{"notImplemented": true})
}

func (b *BlogsHandler) GetBlog(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"notImplemented": true})
}

func (b *BlogsHandler) UpdateBlog(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"notImplemented": true})
}

func (b *BlogsHandler) DeleteBlog(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"notImplemented": true})
}
