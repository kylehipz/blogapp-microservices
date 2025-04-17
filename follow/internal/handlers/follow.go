package handlers

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/follow/internal/services"
)

type FollowHandler struct {
	FollowService services.FollowService
}

func (f *FollowHandler) FollowUser(c echo.Context) error {
	log.Println("FOLLOW USER")
	return nil
}

func (f *FollowHandler) UnfollowUser(c echo.Context) error {
	log.Println("UNFOLLOW USER")
	return nil
}
