package main

import (
	"os"
	"thegraduate-server/auth"
	"thegraduate-server/config"
	"thegraduate-server/graduate_form"
	"thegraduate-server/helper"
	"thegraduate-server/lecturer"
	"thegraduate-server/model"
	"thegraduate-server/redis"
	"thegraduate-server/semester"
	"thegraduate-server/students_description"
	"thegraduate-server/users"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func InitEcho() *echo.Echo {

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	return app
}

func main() {



	app := fx.New(fx.Provide(InitEcho),
		auth.AuthModule,
		users.UserModule,
		semester.SemesterModule,
		students_description.StudentDesc,
		graduate_form.GraduateModule,
		lecturer.LecturerModule,
		fx.Provide(config.InitDatabase),
		fx.Provide(redis.NewRedisClient),
		fx.Invoke(func(e *echo.Echo) {

			e.Use(middleware.Recover())
			// e.Logger.SetOutput(io.Discard)

			// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			// 	Output: io.Discard,
			// }))
			e.GET("/", func(ctx echo.Context) error {

				return ctx.JSON(200, model.ResponseModel{
					Data:    []string{},
					Message: "welcome to the graduate API",
					Status:  "success",
				})
			})
			e.HTTPErrorHandler = helper.ErrorHandler

			var port string = os.Getenv("PORT")

			if port == "" {
				port = "3000"
			}

			var hostPort string = "0.0.0.0:" + port
			e.Start(hostPort)
		}))

	app.Run()

}
