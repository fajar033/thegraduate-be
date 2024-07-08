package students_description

import (
	"context"
	"thegraduate-server/entities"
	"thegraduate-server/helper"
	"thegraduate-server/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type studentDescHandler struct {
	usecase    IStudentDescUsecase
	repository IStudentDescRepository
}

type IStudentDescHandler interface {
	GetStatisticStudent(e echo.Context) error
	GetAllStudent(e echo.Context) error
	CreateStudentDesc(e echo.Context) error
	FindDocsByStudentId(e echo.Context) error
	FindStudentById(e echo.Context) error
	UploadIdCard(e echo.Context) error
	UploadStudentCard(e echo.Context) error
	UploadThesisFile(e echo.Context) error
	UploadFamilyCard(e echo.Context) error
	UploadGradCertificate(e echo.Context) error
	UploadTempGradCertificate(e echo.Context) error
	UploadPhoto(e echo.Context) error
	UploadValiditySheet(e echo.Context) error
	UploadToeicCertificate(e echo.Context) error
	UploadBirthCertificate(e echo.Context) error
	UploadArticle(e echo.Context) error
	UploadCompetencyCertificate(e echo.Context) error
	UpdateDescription(e echo.Context) error
}

func NewStudentDescHandler(usecase IStudentDescUsecase, repository IStudentDescRepository) IStudentDescHandler {
	return &studentDescHandler{
		usecase:    usecase,
		repository: repository,
	}
}

func (u *studentDescHandler) GetStatisticStudent(e echo.Context) error {
	data := u.repository.GetStatisticStudent(context.Background())

	return e.JSON(200, model.ResponseModel{
		Data:    data,
		Message: "success retrieved statistic student",
		Status:  "success",
	})

}

func (u *studentDescHandler) FindDocsByStudentId(e echo.Context) error {

	var studentId string = e.Param("studentId")

	result := u.usecase.FindDocsByStudentId(context.Background(), studentId)
	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Status:  "success",
		Message: "success retrieved document",
	})

}

func (u *studentDescHandler) UploadPhoto(e echo.Context) error {

	file, err := e.FormFile("photo")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

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

	u.usecase.UploadDocs(context, entities.Document{
		Photo: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadGradCertificate(e echo.Context) error {

	file, err := e.FormFile("grad_certificate")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		GradCertificate: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadArticle(e echo.Context) error {

	file, err := e.FormFile("article")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		Article: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadCompetencyCertificate(e echo.Context) error {

	file, err := e.FormFile("competency_certificate")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		CompetencyCertificate: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadTempGradCertificate(e echo.Context) error {

	file, err := e.FormFile("temp_grad")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		TempGradCertificate: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadValiditySheet(e echo.Context) error {

	file, err := e.FormFile("validity_sheet")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		ValiditySheet: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadFamilyCard(e echo.Context) error {

	file, err := e.FormFile("family_card")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		FamilyCard: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadThesisFile(e echo.Context) error {

	file, err := e.FormFile("thesis")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		ThesisFile: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadStudentCard(e echo.Context) error {

	file, err := e.FormFile("student_card")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		StudentCard: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadIdCard(e echo.Context) error {

	file, err := e.FormFile("idcard")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		IdCard: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data",
	})
}

func (u *studentDescHandler) UploadToeicCertificate(e echo.Context) error {

	file, err := e.FormFile("toeic")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		Toeic: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data ",
	})
}

func (u *studentDescHandler) UploadBirthCertificate(e echo.Context) error {

	file, err := e.FormFile("birth_certificate")

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	if err != nil {
		panic(model.BadRequestError{Message: "invalid payload requested"})
	}
	source, err := file.Open()

	if err != nil {
		panic(&model.BadRequestError{Message: "invalid while opening file"})
	}

	defer source.Close()

	if !helper.ValidatePdf(file.Filename) {
		panic(&model.BadRequestError{"file must be a pdf"})
	}

	context := context.Background()

	var fileName string = helper.UploadFile(context, source)

	u.usecase.UploadDocs(context, entities.Document{
		BirthCertificate: fileName,
	}, claims.Email)

	return e.JSON(200, model.ResponseModel{
		Status:  "success",
		Data:    fileName,
		Message: "success upload data ",
	})
}

func (u *studentDescHandler) CreateStudentDesc(e echo.Context) error {

	var payload entities.Student

	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: err.Error()})
	}

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtClaimsCustom)

	payload.FirstName = claims.FirstName
	payload.LastName = claims.LastName
	payload.Email = claims.Email
	u.usecase.Insert(context.Background(), payload)

	return e.JSON(200, model.ResponseModel{
		Data:    payload,
		Message: "students description has been created successfully",
		Status:  "success",
	})
}

func (u *studentDescHandler) GetAllStudent(e echo.Context) error {

	academic := e.QueryParam("academic_year")
	semester := e.QueryParam("semester")
	verified := e.QueryParam("verified")
	Name := e.QueryParam("name")
	verifiedSKL := e.QueryParam("verified_skl")

	result, _ := u.usecase.FindAll(context.Background(), model.QueryStudentDesc{
		AcademicYear: academic,
		Semester:     semester,
		Name:         Name,
		Verified:     verified,
		VerifiedSKL:  verifiedSKL,
	})
	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "succes retreived data",
		Status:  "succes",
	})
}

func (u *studentDescHandler) FindStudentById(e echo.Context) error {
	var id = e.Param("studentId")

	result := u.usecase.FindById(id, context.Background())

	return e.JSON(200, model.ResponseModel{
		Data:    result,
		Message: "successfully fetching a students description",
		Status:  "success",
	})
}

func (u *studentDescHandler) UpdateDescription(e echo.Context) error {
	//TODO implement me
	var id = e.Param("id")
	var payload entities.Student

	if err := e.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: err.Error()})
	}
	u.usecase.UpdateOne(context.Background(), payload, id)

	return e.JSON(200, model.ResponseModel{
		Data:    payload,
		Message: "success update data",
		Status:  "success",
	})
}
