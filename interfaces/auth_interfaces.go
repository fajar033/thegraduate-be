package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"thegraduate-server/model"
)

type IAuthHandler interface {
	Login(c echo.Context) error
}

type IAuthUsecase interface {
	GenerateToken(ctx context.Context, model model.LoginModel) string
}
