package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"thegraduate-server/entities"
	"thegraduate-server/model"
)

func JwtConfigEcho(c echo.Context) jwt.Claims {
	return new(entities.JwtClaimsCustom)
}

func SigningKey() string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
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
