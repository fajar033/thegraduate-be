package config

import (
	"net/http"
	"os"
	"thegraduate-server/entities"
	"thegraduate-server/model"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtConfigEcho(c echo.Context) jwt.Claims {
	return new(entities.JwtClaimsCustom)
}

func SigningKey() string {

	return os.Getenv("JWT_SECRET_KEY")
}

var ConfigJwt echojwt.Config = echojwt.Config{NewClaimsFunc: JwtConfigEcho,
	SigningKey: []byte(SigningKey()),
	ErrorHandler: func(c echo.Context, err error) error {
		return c.JSON(http.StatusBadRequest, model.ResponseModelFailed{
			Message: "Auth Error: " + err.Error(),
			Status:  "failed",
		})
	}}
