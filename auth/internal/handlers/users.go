package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
)

type UsersHandler struct {
	UsersService services.UsersService
}

type RegisterUserRequestBody struct {
	Username string
	Password string
	Email    string
}

func (u *UsersHandler) RegisterUser(c echo.Context) error {
	body := new(RegisterUserRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdUser, err := u.UsersService.CreateUser(
		context.TODO(),
		body.Username,
		body.Email,
		body.Password,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (u *UsersHandler) LoginUser(c echo.Context) error {
	fmt.Println("LOGIN USER!")
	return nil
}
