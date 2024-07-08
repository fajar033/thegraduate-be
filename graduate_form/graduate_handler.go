package graduate_form

import (
	"thegraduate-server/entities"
	"thegraduate-server/model"

	"github.com/labstack/echo/v4"
)

type GraduateFormHandler struct {
	repository GraduateFormRepository
}

func NewGraduateHandler(repository GraduateFormRepository) GraduateFormHandler {

	return GraduateFormHandler{repository}
}

func (g *GraduateFormHandler) GraduateFormHandlerUpload(e echo.Context) error {

	var payload entities.GraduateForm
	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: err.Error()})
	}

	g.repository.Upload(payload)

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success save form graduate",
		Status:  "success",
	})
}

func (g *GraduateFormHandler) Update(e echo.Context) error {
	var payload entities.GraduateForm
	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: err.Error()})
	}
	g.repository.Update(payload.StudentID, payload)
	return e.JSON(200, model.ResponseModel{
		Data:    "-",
		Status:  "success",
		Message: "success updated data",
	})
}

func (g *GraduateFormHandler) GraduateFormHandlerFindById(e echo.Context) error {
	id := e.Param("id")

	result, err := g.repository.FindById(id)

	if err != nil {
		panic(err)
	}

	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "success retrieved data",
		Status:  "success",
	})

}
