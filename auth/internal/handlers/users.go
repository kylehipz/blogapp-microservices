package handlers

import (
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

type LoginUserRequestBody struct {
	Username string
	Password string
}

func (u *UsersHandler) RegisterUser(c echo.Context) error {
	body := new(RegisterUserRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdUser, err := u.UsersService.CreateUser(
		c.Request().Context(),
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
	body := new(LoginUserRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := u.UsersService.Login(c.Request().Context(), body.Username, body.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"accessToken": accessToken})
}
