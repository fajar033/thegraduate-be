package users

import (
	"thegraduate-server/config"
	"thegraduate-server/interfaces"
	"thegraduate-server/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RegisterStaffHandler(echo *echo.Echo, handler interfaces.IUserHandle) {

	echo.POST("/users", handler.Create)
	echo.POST("/users/changepass", handler.ChangePassword)
	echo.POST("/users/sendemail", handler.SendEmail)
	echo.GET("/users", handler.List, echojwt.WithConfig(config.ConfigJwt), middlewares.AdminMiddleware)
	echo.GET("/users/:id", handler.FindById, echojwt.WithConfig(config.ConfigJwt))

}

var UserModule fx.Option = fx.Options(
	fx.Provide(NewUserHandler),
	fx.Provide(NewUserUseCase),
	fx.Provide(NewStaffRepository),
	fx.Invoke(RegisterStaffHandler),
)
