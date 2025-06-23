package handlers

import (
	"net/http"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/errs"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/auth/internal/services"
)

type UsersHandler struct {
	usersService *services.UsersService
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

func NewUsersHandler(usersService *services.UsersService) *UsersHandler {
	return &UsersHandler{usersService: usersService}
}

func (u *UsersHandler) RegisterUser(c echo.Context) error {
	body := new(RegisterUserRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	createdUser, err := u.usersService.CreateUser(
		c.Request().Context(),
		body.Username,
		body.Email,
		body.Password,
	)
	if err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (u *UsersHandler) LoginUser(c echo.Context) error {
	body := new(LoginUserRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	accessToken, err := u.usersService.Login(c.Request().Context(), body.Username, body.Password)
	if err != nil {
		return echo.NewHTTPError(errs.GetHttpStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"accessToken": accessToken})
}

func (u *UsersHandler) TestRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
