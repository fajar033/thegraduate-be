package students_description

import (
	"context"
	"errors"
	"fmt"
	"thegraduate-server/entities"
	"thegraduate-server/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type studentUseCase struct {
	repo IStudentDescRepository
}

type IStudentDescUsecase interface {
	FindAll(ctx context.Context, query model.QueryStudentDesc) ([]entities.DocumentAndStudent, error)
	FindById(id string, ctx context.Context) entities.Student
	Insert(context.Context, entities.Student) entities.Student
	UpdateOne(ctx context.Context, student entities.Student, id string)
	FindDocsByStudentId(context context.Context, id string) entities.Document
	UploadDocs(ctx context.Context, data entities.Document, email string)
}

func NewStudentDescUseCase(repo IStudentDescRepository) IStudentDescUsecase {

	return &studentUseCase{
		repo: repo,
	}
}

func (u *studentUseCase) FindDocsByStudentId(context context.Context, id string) entities.Document {

	result, err := u.repo.FindDocsByStudentId(context, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Document{}
		}
		panic(err)
	}
	const noFileUploaded = "No file uploaded for this field"

	if result.BirthCertificate != "" {
		result.BirthCertificate = "https://storage.googleapis.com/thegraduate-bucket/" + result.BirthCertificate
	} else {
		result.BirthCertificate = noFileUploaded
	}

	if result.IdCard != "" {
		result.IdCard = "https://storage.googleapis.com/thegraduate-bucket/" + result.IdCard
	} else {
		result.IdCard = noFileUploaded
	}

	if result.Toeic != "" {
		result.Toeic = "https://storage.googleapis.com/thegraduate-bucket/" + result.Toeic
	} else {
		result.Toeic = noFileUploaded
	}

	if result.Article != "" {
		result.Article = "https://storage.googleapis.com/thegraduate-bucket/" + result.Article
	} else {
		result.Article = noFileUploaded
	}

	if result.StudentCard != "" {
		result.StudentCard = "https://storage.googleapis.com/thegraduate-bucket/" + result.StudentCard
	} else {
		result.StudentCard = noFileUploaded
	}

	if result.GradCertificate != "" {
		result.GradCertificate = "https://storage.googleapis.com/thegraduate-bucket/" + result.GradCertificate
	} else {
		result.GradCertificate = noFileUploaded
	}

	if result.TempGradCertificate != "" {
		result.TempGradCertificate = "https://storage.googleapis.com/thegraduate-bucket/" + result.TempGradCertificate
	} else {
		result.TempGradCertificate = noFileUploaded
	}

	if result.CompetencyCertificate != "" {
		result.CompetencyCertificate = "https://storage.googleapis.com/thegraduate-bucket/" + result.CompetencyCertificate
	} else {
		result.CompetencyCertificate = noFileUploaded
	}

	if result.ThesisFile != "" {
		result.ThesisFile = "https://storage.googleapis.com/thegraduate-bucket/" + result.ThesisFile
	} else {
		result.ThesisFile = noFileUploaded
	}

	if result.ValiditySheet != "" {
		result.ValiditySheet = "https://storage.googleapis.com/thegraduate-bucket/" + result.ValiditySheet
	} else {
		result.ValiditySheet = noFileUploaded
	}

	if result.FamilyCard != "" {
		result.FamilyCard = "https://storage.googleapis.com/thegraduate-bucket/" + result.FamilyCard
	} else {
		result.FamilyCard = noFileUploaded
	}

	if result.Photo != "" {
		result.Photo = "https://storage.googleapis.com/thegraduate-bucket/" + result.Photo
	} else {
		result.Photo = noFileUploaded
	}
	return result

}

func (u *studentUseCase) UploadDocs(ctx context.Context, data entities.Document, email string) {

	result, _ := u.repo.FindByEmail(ctx, email)
	_, err := u.repo.FindDocsByStudentId(ctx, result.StudentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			data.Id = uuid.New().String()
			data.StudentId = result.StudentID

			err = u.repo.UploadDocs(ctx, data, false)

			if err != nil {

				panic(err)
			}
		} else {
			panic(err)
		}
	}
	data.StudentId = result.StudentID
	err = u.repo.UploadDocs(ctx, data, true)

	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			panic(&model.NotFoundError{"no related data found"})
		}

		panic(err)
	}

}

func (u *studentUseCase) FindAll(ctx context.Context, query model.QueryStudentDesc) ([]entities.DocumentAndStudent, error) {
	fmt.Print(query)
	result, err := u.repo.FindAll(ctx, query)

	var resultDocument []entities.DocumentAndStudent
	for _, data := range result {
		resultDocumentDB := u.FindDocsByStudentId(ctx, data.StudentID)
		resultDocument = append(resultDocument, entities.DocumentAndStudent{
			Data:     data,
			Document: resultDocumentDB,
		})
	}

	if err != nil {
		panic(err)
	}
	return resultDocument, nil

}

func (u *studentUseCase) Insert(context context.Context, data entities.Student) entities.Student {

	result, err := u.repo.FindOneById(context, data.StudentID)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.BadRequestError{Message: err.Error()})
		}
	}

	if result.StudentID != "" {
		panic(&model.ConflictError{Message: "students description already created"})
	}
	data.Verification = "NOT_REGISTERED"
	data.VerificationSKL = "NOT_VERIFIED"
	data.Message = "NO MESSAGE AVAILABLE"
	data.MessageSKL = "NO MESSAGE AVAILABLE"
	err = u.repo.Create(data, context)
	if err != nil {
		panic(err)
	}
	return data
}

func (u *studentUseCase) FindById(id string, ctx context.Context) entities.Student {

	result, err := u.repo.FindOneById(ctx, id)

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(model.NotFoundError{Message: "no data found"})
		}
		panic(err)
	}

	return result

}

func (u *studentUseCase) UpdateOne(ctx context.Context, student entities.Student, id string) {
	_, err := u.repo.FindOneById(ctx, id)

	if err != nil {

		if err != gorm.ErrRecordNotFound {
			panic(&model.NotFoundError{Message: "no data found"})
		}
		panic(err)
	}

	err = u.repo.Update(ctx, student, id)
	if err != nil {
		panic(err)
	}

}
