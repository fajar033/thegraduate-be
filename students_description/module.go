package students_description

import (
	"thegraduate-server/config"
	"thegraduate-server/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RegisterStudentDescModule(echo *echo.Echo, handler IStudentDescHandler) {
	echo.GET("/students/docs/:studentId", handler.FindDocsByStudentId)
	echo.POST("/students/doc/studentcard", handler.UploadStudentCard, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/valsheet", handler.UploadValiditySheet, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/thesis", handler.UploadThesisFile, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/gradcert", handler.UploadGradCertificate, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/tempgrad", handler.UploadTempGradCertificate, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/familycard", handler.UploadFamilyCard, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/idcard", handler.UploadIdCard, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/photo", handler.UploadPhoto, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/toeic", handler.UploadToeicCertificate, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/birth", handler.UploadBirthCertificate, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/article", handler.UploadArticle, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students/doc/comp_cert", handler.UploadCompetencyCertificate, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.GET("/students", handler.GetAllStudent, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.POST("/students", handler.CreateStudentDesc, echojwt.WithConfig(config.ConfigJwt), middlewares.UserMiddleware)
	echo.GET("/students/:studentId", handler.FindStudentById)
	echo.PATCH("/students/:id"		, handler.UpdateDescription)
	echo.GET("/students/statistic", handler.GetStatisticStudent)
}

var StudentDesc = fx.Options(fx.Invoke(RegisterStudentDescModule),
	fx.Provide(NewStudentDescUseCase),
	fx.Provide(NewStudentDescRepository),
	fx.Provide(NewStudentDescHandler))
