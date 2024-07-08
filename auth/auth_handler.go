package auth

import (
	"context"
	"github.com/labstack/echo/v4"
	"thegraduate-server/interfaces"
	"thegraduate-server/model"
)

type authHandler struct {
	AuthUsecase interfaces.IAuthUsecase
}

func NewAuthHandler(usecase interfaces.IAuthUsecase) interfaces.IAuthHandler {

	return &authHandler{AuthUsecase: usecase}
}

func (a authHandler) Login(c echo.Context) error {

	var payload model.LoginModel
	if err := c.Bind(&payload); err != nil {
		panic(&model.BadRequestError{Message: "invalid payload requested"})
	}

	var token string = a.AuthUsecase.GenerateToken(context.Background(), payload)

	return c.JSON(200, model.ResponseModel{
		Data:    token,
		Message: "success created token",
		Status:  "success",
	})
}
