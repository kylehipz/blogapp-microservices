package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/follow/internal/services"
)

type FollowHandler struct {
	FollowService *services.FollowService
}

type FollowUserRequestBody struct {
	Followee string `json:"followee"`
}

func (f *FollowHandler) FollowUser(c echo.Context) error {
	// parse user from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.JwtCustomClaims)
	follower := claims.ID

	// parse request body
	body := new(FollowUserRequestBody)

	createdFollow, err := f.FollowService.FollowUser(c.Request().Context(), follower, body.Followee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, createdFollow)
}

func (f *FollowHandler) UnfollowUser(c echo.Context) error {
	// parse user from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.JwtCustomClaims)
	follower := claims.ID

	// parse request body
	body := new(FollowUserRequestBody)

	err := f.FollowService.UnfollowUser(c.Request().Context(), follower, body.Followee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
