package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"thegraduate-server/entities"
	"thegraduate-server/model"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*entities.JwtClaimsCustom)

		if claims.Role == "user" {
			return c.JSON(401, model.ResponseModelFailed{
				Message: "cannot access the endpoint",
				Status:  "failed",
			})
		}

		return next(c)
	}
}
