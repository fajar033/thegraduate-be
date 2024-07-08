package users

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"os"
	"thegraduate-server/entities"
	"thegraduate-server/helper"
	"thegraduate-server/interfaces"
	"thegraduate-server/model"
	"thegraduate-server/redis"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo interfaces.IUserRepository
	redis    *redis.RedisClient
}

//go:embed mail.htm
var templates embed.FS
var myTemplates = template.Must(template.ParseFS(templates, "mail.htm"))

func NewUserUseCase(userRepo interfaces.IUserRepository, redis *redis.RedisClient) interfaces.IUserUseCase {

	return &userUsecase{userRepo: userRepo, redis: redis}
}

func (u *userUsecase) FindAll(ctx context.Context) []entities.UserEntity {
	result := u.userRepo.FindAll(ctx)

	return result
}

func (u *userUsecase) Create(ctx context.Context, payload model.UserModel) entities.UserEntity {

	message := helper.Validate[model.UserModel](payload)

	if message != nil {

		panic(&model.ValidationError{ErrMessage: message})
	}

	_, data := u.userRepo.FindOne(ctx, payload.Username)

	if data != (entities.UserEntity{}) {
		panic(&model.ConflictError{Message: "users is already created"})
	}
	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(&model.BadRequestError{Message: "error while generate hash password"})
	}

	var result = entities.UserEntity{
		Id:        payload.Id,
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  string(hashedPassword),
		Role:      payload.Role,
		Email:     payload.Email,
	}
	err = u.userRepo.Insert(ctx, result)

	if err != nil {
		panic(err)
	}

	return result

}

func (u *userUsecase) SendEmail(ctx context.Context, email string) {
	var body bytes.Buffer

	user, err := u.userRepo.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: err.Error()})
		}
		panic(err)
	}

	session, _ := gonanoid.New(10)

	var bodyTemplate map[string]any = map[string]any{
		"username": user.Username,
		"email":    user.Email,
		"session":  os.Getenv("URL_FE") + fmt.Sprintf("?token=%s&username=%s", session, user.Username),
	}

	resultUser, err := u.userRepo.FindOneSessionByUsername(ctx, user.Username)

	if resultUser.Username != "" {
		panic(&model.ConflictError{Message: "token already exists"})
	}

	err = u.userRepo.InsertSession(ctx, entities.SessionEntity{
		SessionCode: session,
		Username:    user.Username,
	})

	if err != nil {
		panic(err)
	}

	err = myTemplates.Execute(&body, bodyTemplate)

	if err != nil {
		panic(err.Error())
	}
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "The Graduate")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func (u *userUsecase) ForgetPassword(ctx context.Context, data model.Forgetpassword) {
	//TODO implement me
	result, err := u.userRepo.FindOneSessionByUsername(ctx, data.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{err.Error()})
		}
		panic(err)
	}
	if result.SessionCode != data.Session {
		panic(&model.BadRequestError{Message: "Invalid session token"})
	}
	fmt.Print("kesini")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(&model.BadRequestError{Message: "error while generate hash password"})
	}
	err = u.userRepo.ChangePassword(ctx, string(hashedPassword), data.Username)

	if err != nil {
		panic(err)
	}
	err = u.userRepo.DeleteSession(ctx, data.Username)
	if err != nil {
		panic(err)
	}

}
