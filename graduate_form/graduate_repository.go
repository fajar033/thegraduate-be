package graduate_form

import (
	"context"
	"errors"
	"fmt"
	"thegraduate-server/entities"
	"thegraduate-server/model"
	"thegraduate-server/students_description"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type GraduateFormRepository struct {
	db         *gorm.DB
	repository students_description.IStudentDescRepository
}

func NewGraduateRepository(db *gorm.DB, repo students_description.IStudentDescRepository) GraduateFormRepository {

	return GraduateFormRepository{db: db,
		repository: repo}
}

func (c GraduateFormRepository) Upload(form entities.GraduateForm) {
	id, err := gonanoid.New(10)
	fmt.Print(form)
	form.Id = id

	err = c.db.WithContext(context.Background()).Table("graduate_certificate_form").Create(&form).Error

	if err != nil {
		panic(err)
	}
}
func (c GraduateFormRepository) FindById(id string) (*entities.GraduateForm, error) {

	var result entities.GraduateForm
	err := c.db.WithContext(context.Background()).Table("graduate_certificate_form").Where("student_id = ?", id).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: "no record found"})
		}
		panic(err)
	}

	return &result, nil

}
func (c GraduateFormRepository) Update(id string, form entities.GraduateForm) {

	_, err := c.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{"no record found"})
		}
		panic(err)
	}

	err = c.db.WithContext(context.Background()).Table("graduate_certificate_form").Where("student_id=?", id).Updates(&form).Error

	if err != nil {
		panic(err)

	}
}
