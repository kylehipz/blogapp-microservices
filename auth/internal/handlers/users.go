package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
)

type UsersHandler struct {
	UsersService services.UsersService
}

func (u *UsersHandler) RegisterUser(c echo.Context) error {
	fmt.Println("REGISTER USER!")
	return nil
}

func (u *UsersHandler) LoginUser(c echo.Context) error {
	fmt.Println("LOGIN USER!")
	return nil
}
