package graduate_form

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewGraduateFormRegistry(e *echo.Echo, handler GraduateFormHandler) {
	e.GET("/graduateform/detail/:id", handler.GraduateFormHandlerFindById)
	e.POST("/graduateform/create", handler.GraduateFormHandlerUpload)
	e.PATCH("/graduateform/update", handler.Update)
}

var GraduateModule fx.Option = fx.Options(
	fx.Invoke(NewGraduateFormRegistry),
	fx.Provide(NewGraduateHandler),
	fx.Provide(NewGraduateRepository))
