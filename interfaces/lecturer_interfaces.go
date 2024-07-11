package interfaces

import (
	"context"
	"thegraduate-server/entities"

	"github.com/labstack/echo/v4"
)

type ILecturerHandler interface {
	GetStaticLecturer(e echo.Context) error
	List(e echo.Context) error
	Create(e echo.Context) error
	Update(e echo.Context) error
	// FindByInitialLecturer(e echo.Context) error
	GetSKLByNidnAndStudentId(e echo.Context) error
	Delete(e echo.Context) error
	UploadAdvAssignmentLetter(e echo.Context) error
	UploadInvitation(e echo.Context) error
	UploadTempGrad(e echo.Context) error
	UploadExaminerAssignmentLetter(e echo.Context) error
	UploadDocsOfficialReport(e echo.Context) error
	FindAllDocuments(e echo.Context) error
	FindOneLecturer(e echo.Context) error
}

type ILecturerRepository interface {
	FindOneLecturerById(ctx context.Context, id string) (*entities.Lecturer, error)
	GetTotalStatic(ctx context.Context, nidn string) map[string]any
	FindDocumentByStudentAndNidn(ctx context.Context, studentId string, nidn string) (entities.DocumentAdminEntity, error)
	FindOneById(ctx context.Context, id string) (map[string]any, error)
	Insert(ctx context.Context, data entities.Lecturer)
	Update(ctx context.Context, data entities.Lecturer, nidn string)
	Delete(ctx context.Context, id string)
	FindAllDocuments(ctx context.Context, nidn string, name string) []entities.StudentDocumentAdmin
	UploadDocsAdmin(ctx context.Context, data entities.DocumentAdminEntity, isDocumentAlreadyCreated bool)
	List(ctx context.Context) []map[string]interface{}
}
