package users

import (
	"context"
	"thegraduate-server/interfaces"
	"thegraduate-server/model"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase interfaces.IUserUseCase
}

func NewUserHandler(usecase interfaces.IUserUseCase) interfaces.IUserHandle {
	return &UserHandler{usecase: usecase}
}

func (u *UserHandler) ChangePassword(ctx echo.Context) error {
	//TODO implement me

	var payload model.Forgetpassword

	if err := ctx.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}

	u.usecase.ForgetPassword(context.Background(), payload)

	return ctx.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success update password",
		Status:  "status",
	})

}

func (u *UserHandler) List(ctx echo.Context) error {

	result := u.usecase.FindAll(context.Background())

	return ctx.JSON(200, model.ResponseModel{
		Status: "success",
		Data:   result,
	})

}

func (u *UserHandler) Create(ctx echo.Context) error {

	var payload model.UserModel

	if err := ctx.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}

	id := u.usecase.Create(context.Background(), payload)

	return ctx.JSON(201, model.ResponseModel{
		Data:    id,
		Message: "successfully created users",
		Status:  "success",
	})

}

func (u *UserHandler) FindById(ctx echo.Context) error {

	id := ctx.Param("id")
	return ctx.JSON(200, model.ResponseModel{
		Data:    id,
		Message: "successfully retrieved data students_description",
		Status:  "success",
	})
}

func (u *UserHandler) SendEmail(ctx echo.Context) error {

	var payload model.EmailModel

	if err := ctx.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}

	u.usecase.SendEmail(context.Background(), payload.Email)

	return ctx.JSON(200, model.ResponseModel{
		Data:    nil,
		Message: "success send an email",
		Status:  "success",
	})

}

func (u *UserHandler) UploadFiles(ctx echo.Context) error {
	panic("w")
}
