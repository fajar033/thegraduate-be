package semester

import (
	"strconv"
	"thegraduate-server/model"

	"github.com/labstack/echo/v4"
)

type SemesterHandler struct {
	repo SemesterRepository
}

func NewHandlerSemester(repo SemesterRepository) SemesterHandler {
	return SemesterHandler{repo: repo}
}

func (s *SemesterHandler) Delete(ctx echo.Context) error {

	id := ctx.Param("id")

	s.repo.Delete(id)

	return ctx.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "succeess deleted data",
		Status:  "success",
	})
}

func (s *SemesterHandler) FindById(e echo.Context) error {

	id := e.Param("id")
	result := s.repo.FindById(id)

	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "success find a data",
		Status:  "success",
	})

}

func (s *SemesterHandler) Create(e echo.Context) error {

	var payload model.CreateSemesterModel

	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}

	data, err := strconv.Atoi(payload.AcademicYear)

	if err != nil {
		panic(&model.BadRequestError{err.Error()})
	}
	payload.AcademicYear = payload.AcademicYear + "/" + strconv.Itoa(data+1)

	s.repo.CreateSemester(payload)

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "succeess created data",
		Status:  "success",
	})
}

func (s *SemesterHandler) List(e echo.Context) error {

	result := s.repo.GetSemesters()

	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "success get list semester",
		Status:  "success",
	})

}
