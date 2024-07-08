package lecturer

import (
	"thegraduate-server/interfaces"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RegisterRouteLecturer(handler interfaces.ILecturerHandler, echo *echo.Echo) {

	echo.POST("/lecturer/create", handler.Create)
	echo.PATCH("/lecturer/update/:nidn", handler.Update)
	echo.DELETE("/lecturer/:nidn", handler.Delete)
	echo.GET("/lecturer/list", handler.List)
	echo.GET("/lecturer/detail/:id", handler.FindOneLecturer)
	echo.GET("/lecturer/skl/:studentid/:nidn", handler.GetSKLByNidnAndStudentId)
	echo.GET("/lecturer/document/list", handler.FindAllDocuments)
	echo.GET("/lecturer/statistic", handler.GetStaticLecturer)
	echo.PATCH("/admin/docs/offreport", handler.UploadDocsOfficialReport)
	echo.PATCH("/admin/docs/examinerletter", handler.UploadExaminerAssignmentLetter)
	echo.PATCH("/admin/docs/advisorletter", handler.UploadAdvAssignmentLetter)
	echo.PATCH("/admin/docs/invitation", handler.UploadInvitation)
	echo.PATCH("/admin/docs/tempgrad", handler.UploadTempGrad)
}

var LecturerModule fx.Option = fx.Options(
	fx.Provide(NewLecturerService),
	fx.Provide(NewLecturerHandler),
	fx.Invoke(RegisterRouteLecturer),
	fx.Provide(NewLecturerRepository),
)
