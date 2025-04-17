package handlers

import (
	"net/http"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/blogs/internal/services"
)

type BlogsHandler struct {
	BlogsService *services.BlogsService
}

type CreateBlogRequestBody struct {
	Content string `json:"content"`
}

func (b *BlogsHandler) CreateBlog(c echo.Context) error {
	body := new(CreateBlogRequestBody)
	author := middlewares.GetUserID(c)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdBlog, err := b.BlogsService.CreateBlog(c.Request().Context(), author, body.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, createdBlog)
}

func (b *BlogsHandler) GetBlog(c echo.Context) error {
	blogID := c.Param("id")

	blog, err := b.BlogsService.GetBlog(c.Request().Context(), blogID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, blog)
}

func (b *BlogsHandler) UpdateBlog(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"notImplemented": true})
}

func (b *BlogsHandler) DeleteBlog(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"notImplemented": true})
}
