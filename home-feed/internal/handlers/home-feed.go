package handlers

import (
	"net/http"
	"strconv"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/errs"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/middlewares"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal/services"
)

type HomeFeedHandler struct {
	homeFeedService *services.HomeFeedService
	logger          *zap.Logger
}

func NewHomeFeedHandler(
	homeFeedService *services.HomeFeedService,
	logger *zap.Logger,
) *HomeFeedHandler {
	return &HomeFeedHandler{homeFeedService: homeFeedService, logger: logger}
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

	homeFeed, err := h.homeFeedService.GetHomeFeed(
		c.Request().Context(),
		userID,
		createdAt,
		int32(limit),
	)
	if err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	if homeFeed == nil {
		homeFeed = []*types.Blog{}
	}

	return c.JSON(http.StatusOK, homeFeed)
}
