package semester

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RegisterSemester(e *echo.Echo, handler SemesterHandler) {
	e.GET("/semester/list", handler.List)
	e.POST("/semester/create", handler.Create)
	e.DELETE("/semester/:id", handler.Delete)
	e.GET("/semester/detail/:id", handler.List)

}

var SemesterModule = fx.Options(
	fx.Provide(NewSemesterRepository),
	fx.Provide(NewHandlerSemester),
	fx.Invoke(RegisterSemester),
)
