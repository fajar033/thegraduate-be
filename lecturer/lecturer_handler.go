package lecturer

import (
	"context"
	"thegraduate-server/entities"
	"thegraduate-server/helper"
	"thegraduate-server/interfaces"
	"thegraduate-server/model"

	"github.com/labstack/echo/v4"
)

type lecturerHandler struct {
	repo    interfaces.ILecturerRepository
	service ILecturerService
}

func NewLecturerHandler(repo interfaces.ILecturerRepository, service ILecturerService) interfaces.ILecturerHandler {
	return &lecturerHandler{
		repo:    repo,
		service: service,
	}
}

func (l *lecturerHandler) FindOneLecturer(e echo.Context) error {

	id := e.Param("id")

	result,_ := l.repo.FindOneLecturerById(context.Background(), id)

	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "success get detail data",
		Status:  "success",
	})
}

func (l *lecturerHandler) GetStaticLecturer(e echo.Context) error {
	//TODO implement me
	var total = l.repo.GetTotalStatic(context.Background())
	return e.JSON(200, model.ResponseModel{
		Data:    total,
		Message: "success retrieved total data",
		Status:  "success",
	})
}

func (l *lecturerHandler) FindAllDocuments(e echo.Context) error {
	nidn := e.QueryParam("nidn")
	name := e.QueryParam("name")
	results := l.repo.FindAllDocuments(context.Background(), nidn, name)

	return e.JSON(200, model.ResponseModel{
		Data:    results,
		Status:  "success",
		Message: "success retrieved document admin",
	})
}

func (l *lecturerHandler) UploadTempGrad(e echo.Context) error {
	//TODO implement me

	file, err := e.FormFile("temp_grad")
	nidn := e.FormValue("nidn")
	npm := e.FormValue("npm")
	if err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)
	l.service.UploadDocs(context, entities.DocumentAdminEntity{
		Nidn:      nidn,
		StudentId: npm,
		TempGrad:  fileName,
	})

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success upload file",
		Status:  "success",
	})

}

func (l *lecturerHandler) UploadAdvAssignmentLetter(e echo.Context) error {
	//TODO implement me

	file, err := e.FormFile("advisor_letter")
	nidn := e.FormValue("nidn")
	npm := e.FormValue("npm")
	if err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)
	l.service.UploadDocs(context, entities.DocumentAdminEntity{
		Nidn:                    nidn,
		StudentId:               npm,
		AdvisorAssignmentLetter: fileName,
	})

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success upload file",
		Status:  "success",
	})

}

func (l *lecturerHandler) UploadInvitation(e echo.Context) error {
	file, err := e.FormFile("invitation")
	nidn := e.FormValue("nidn")
	npm := e.FormValue("npm")
	if err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)
	l.service.UploadDocs(context, entities.DocumentAdminEntity{
		Nidn:       nidn,
		StudentId:  npm,
		Invitation: fileName,
	})

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success upload file",
		Status:  "success",
	})

}

func (l *lecturerHandler) UploadExaminerAssignmentLetter(e echo.Context) error {
	//TODO implement me
	file, err := e.FormFile("examiner_letter")
	nidn := e.FormValue("nidn")
	npm := e.FormValue("npm")
	if err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)
	l.service.UploadDocs(context, entities.DocumentAdminEntity{
		Nidn:                     nidn,
		StudentId:                npm,
		ExaminerAssignmentLetter: fileName,
	})

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success upload file",
		Status:  "success",
	})

}

func (l *lecturerHandler) UploadDocsOfficialReport(e echo.Context) error {

	file, err := e.FormFile("official_report")
	nidn := e.FormValue("nidn")
	npm := e.FormValue("npm")
	if err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)
	l.service.UploadDocs(context, entities.DocumentAdminEntity{
		Nidn:           nidn,
		StudentId:      npm,
		OfficialReport: fileName,
	})

	return e.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success upload file",
		Status:  "success",
	})

}

func (l *lecturerHandler) List(e echo.Context) error {

	result := l.repo.List(context.Background())

	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "success retrieved all data",
		Status:  "success",
	})
}

func (l *lecturerHandler) Create(e echo.Context) error {

	var payload entities.Lecturer

	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: err.Error()})
	}

	l.repo.Insert(context.Background(), payload)

	return e.JSON(201, model.ResponseModel{
		Data:    payload,
		Message: "success update data",
		Status:  "success",
	})
}

func (l *lecturerHandler) Update(e echo.Context) error {

	var payload entities.Lecturer

	var nidn string = e.Param("nidn")

	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: err.Error()})
	}

	l.repo.Update(context.Background(), payload, nidn)

	return e.JSON(201, model.ResponseModel{
		Data:    payload,
		Message: "success add data",
		Status:  "success",
	})

}

func (l *lecturerHandler) GetSKLByNidnAndStudentId(e echo.Context) error {
	result, err := l.repo.FindDocumentByStudentAndNidn(context.Background(), e.Param("studentid"), e.Param("nidn"))

	if err != nil {
		panic(&model.NotFoundError{Message: err.Error()})
	}
	result.AdvisorAssignmentLetter = ""
	result.ExaminerAssignmentLetter = ""

	result.TempGrad = "https://storage.googleapis.com/thegraduate-bucket/" + result.TempGrad
	return e.JSON(200, model.ResponseModel{
		Data:   result,
		Status: "success",

		Message: "success retrieved document temporary graduate certificate",
	})

}

func (l *lecturerHandler) Delete(e echo.Context) error {

	var nidn string = e.Param("nidn")
	l.repo.Delete(context.Background(), nidn)

	return e.JSON(200, model.ResponseModel{
		Data:    nidn,
		Message: "success delete data",
		Status:  "success",
	})

}
