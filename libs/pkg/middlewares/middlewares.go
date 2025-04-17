package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

func NewAuthenticationMiddleware(jwtSecret string) echo.MiddlewareFunc {
	authConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.JwtCustomClaims)
		},
		SigningKey: []byte(jwtSecret),
	}

	return echojwt.WithConfig(authConfig)
}
