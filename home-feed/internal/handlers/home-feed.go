package handlers

import (
	"net/http"
	"strconv"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/services"
)

type HomeFeedHandler struct {
	HomeFeedService *services.HomeFeedService
}

func (h *HomeFeedHandler) GetHomeFeed(c echo.Context) error {
	userID := middlewares.GetUserID(c)
	createdAt := c.QueryParam("createdAt")

	if createdAt == "" {
		createdAt = "now"
	}

	limitStr := c.QueryParam("limit")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	homeFeed, err := h.HomeFeedService.GetHomeFeed(
		c.Request().Context(),
		userID,
		createdAt,
		int32(limit),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if homeFeed == nil {
		homeFeed = []db.Blog{}
	}

	return c.JSON(http.StatusOK, homeFeed)
}
