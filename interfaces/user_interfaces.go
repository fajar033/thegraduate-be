package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"thegraduate-server/entities"
	"thegraduate-server/model"
)

type IUserRepository interface {
	Insert(ctx context.Context, entity entities.UserEntity) error
	FindOne(ctx context.Context, id string) (error, entities.UserEntity)
	FindAll(ctx context.Context) []entities.UserEntity
	FindByEmail(ctx context.Context, email string) (entities.UserEntity, error)
	DeleteSession(ctx context.Context, username string) error
	InsertSession(ctx context.Context, data entities.SessionEntity) error
	ChangePassword(ctx context.Context, password string, username string) error
	FindOneSessionByUsername(ctx context.Context, username string) (entities.SessionEntity, error)
}

type IUserHandle interface {
	SendEmail(ctx echo.Context) error
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	FindById(ctx echo.Context) error
	ChangePassword(ctx echo.Context) error
	UploadFiles(ctx echo.Context) error
}

type IUserUseCase interface {
	ForgetPassword(ctx context.Context, data model.Forgetpassword)
	SendEmail(ctx context.Context, email string)
	FindAll(ctx context.Context) []entities.UserEntity
	Create(ctx context.Context, model model.UserModel) entities.UserEntity
}
