package handlers

import (
	"net/http"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/errs"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/blogs/internal/services"
)

type BlogsHandler struct {
	blogsService *services.BlogsService
}

type CreateBlogRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewBlogsHandler(blogsService *services.BlogsService) *BlogsHandler {
	return &BlogsHandler{blogsService: blogsService}
}

func (b *BlogsHandler) CreateBlog(c echo.Context) error {
	body := new(CreateBlogRequestBody)
	author := middlewares.GetUserID(c)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	createdBlog, err := b.blogsService.CreateBlog(
		c.Request().Context(),
		author,
		body.Title,
		body.Content,
	)
	if err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusCreated, createdBlog)
}

func (b *BlogsHandler) GetBlog(c echo.Context) error {
	blogID := c.Param("id")

	blog, err := b.blogsService.GetBlog(c.Request().Context(), blogID)
	if err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, blog)
}

func (b *BlogsHandler) UpdateBlog(c echo.Context) error {
	blogId := c.Param("id")

	body := new(CreateBlogRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	updateBlog, err := b.blogsService.UpdateBlog(
		c.Request().Context(),
		blogId,
		body.Title,
		body.Content,
	)
	if err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, updateBlog)
}

func (b *BlogsHandler) DeleteBlog(c echo.Context) error {
	blogId := c.Param("id")
	userId := middlewares.GetUserID(c)

	if err := b.blogsService.DeleteBlog(c.Request().Context(), blogId, userId); err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
