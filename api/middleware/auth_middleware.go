package middleware

import (
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{jwtSecret: jwtSecret}
}

func (m *AuthMiddleware) ValidateSecretKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//YP: For the sake of simplicity, we are using a header to pass the secret key
			//YP: it would be better if we use like a JWT Token, to verify and more secure and make sure it's just service to service communication
			SecretKeyHeader := c.Request().Header.Get("X-Secret-Key")

			if SecretKeyHeader != m.jwtSecret {
				return echo.NewHTTPError(401, "Unauthorized")
			}

			return next(c)
		}
	}
}
