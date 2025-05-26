package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

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

func GetUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.JwtCustomClaims)
	userID := claims.ID

	return userID
}

func NewZapLoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			err := next(c)

			stop := time.Now()
			latency := stop.Sub(start)

			logger.Info("HTTP request",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", res.Status),
				zap.Duration("latency", latency),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
				zap.String("request_id", res.Header().Get(echo.HeaderXRequestID)),
				zap.Error(err),
			)

			return nil
		}
	}
}
