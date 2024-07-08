package auth

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"thegraduate-server/interfaces"
)

func RegisterAuthHandler(echo *echo.Echo, handler interfaces.IAuthHandler) {

	echo.POST("/auth/login", handler.Login)

}

var AuthModule fx.Option = fx.Options(
	fx.Invoke(RegisterAuthHandler),
	fx.Provide(NewAuthHandler),
	fx.Provide(NewAuthUsecase))
