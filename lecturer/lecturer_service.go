package lecturer

import (
	"context"
	"errors"
	"thegraduate-server/entities"
	"thegraduate-server/interfaces"
	"thegraduate-server/model"
	"thegraduate-server/students_description"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type ILecturerService interface {
	UploadDocs(ctx context.Context, payload entities.DocumentAdminEntity)
}

type lecturerService struct {
	repo        interfaces.ILecturerRepository
	repoStudent students_description.IStudentDescRepository
}

func NewLecturerService(repo interfaces.ILecturerRepository, student students_description.IStudentDescRepository) ILecturerService {
	return &lecturerService{repo: repo, repoStudent: student}
}

func (l *lecturerService) UploadDocs(ctx context.Context, payload entities.DocumentAdminEntity) {

	_, err := l.repo.FindOneById(ctx, payload.Nidn)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: err.Error()})
		}
		panic(err)
	}

	_, err = l.repoStudent.FindOneById(ctx, payload.StudentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: err.Error()})
		}
		panic(err)
	}

	_, err = l.repo.FindDocumentByStudentAndNidn(ctx, payload.StudentId, payload.Nidn)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			id, _ := gonanoid.New(10)
			payload.Id = id
			l.repo.UploadDocsAdmin(ctx, payload, false)
			return
		}
		panic(err)
	}

	l.repo.UploadDocsAdmin(ctx, payload, true)

}
