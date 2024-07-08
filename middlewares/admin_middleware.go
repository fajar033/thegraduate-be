package middlewares

import (
	"thegraduate-server/entities"
	"thegraduate-server/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*entities.JwtClaimsCustom)

		if claims.Role != "admin" || claims.Role == "lecturer" {
			return c.JSON(401, model.ResponseModelFailed{
				Message: "cannot access the endpoint",
				Status:  "failed",
			})
		}

		return next(c)
	}
}
