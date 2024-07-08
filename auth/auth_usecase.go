package auth

import (
	"context"
	"errors"
	"os"
	"thegraduate-server/entities"
	"thegraduate-server/interfaces"
	model_type "thegraduate-server/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authUsecase struct {
	userRepo interfaces.IUserRepository
}

func NewAuthUsecase(staffRepo interfaces.IUserRepository) interfaces.IAuthUsecase {
	return &authUsecase{userRepo: staffRepo}
}

func (a *authUsecase) GenerateToken(ctx context.Context, model model_type.LoginModel) string {

	err, data := a.userRepo.FindOne(ctx, model.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model_type.NotFoundError{Message: "no data found"})
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(model.Password))

	if err != nil {
		panic(&model_type.NotFoundError{Message: "no data found"})
	}

	var claims *entities.JwtClaimsCustom = &entities.JwtClaimsCustom{
		Npm:       data.Id,
		Username:  data.Username,
		Email:     data.Email,
		Role:      data.Role,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 120)),
		},
	}
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		panic(err)
	}
	return t
}
