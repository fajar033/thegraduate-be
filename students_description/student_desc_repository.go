package students_description

import (
	"context"
	"fmt"
	"thegraduate-server/entities"
	"thegraduate-server/model"

	"gorm.io/gorm"
)

type studentDescRepository struct {
	db *gorm.DB
}

type Pagination struct {
	Page     int
	PageSize int
}

type IStudentDescRepository interface {
	Create(entity entities.Student, ctx context.Context) error
	FindOneById(ctx context.Context, id string) (entities.Student, error)
	FindAll(ctx context.Context, query model.QueryStudentDesc) ([]entities.Student, error)
	UploadDocs(ctx context.Context, data entities.Document, isDocumentAlreadyCreated bool) error
	FindByEmail(ctx context.Context, email string) (entities.Student, error)
	Update(ctx context.Context, student entities.Student, id string) error
	FindDocsByStudentId(ctx context.Context, id string) (entities.Document, error)
	GetStatisticStudent(ctx context.Context) map[string]int64
}

func NewStudentDescRepository(db *gorm.DB) IStudentDescRepository {
	return &studentDescRepository{db: db}
}

func (s *studentDescRepository) GetStatisticStudent(ctx context.Context) map[string]int64 {
	//TODO implement me
	var countVerif int64
	err := s.db.WithContext(ctx).Table("students_description").Where("verification=?", "VERIFIED").Count(&countVerif).Error

	if err != nil {
		panic(err)
	}
	var countNotVerified int64
	err = s.db.WithContext(ctx).Table("students_description").Where("verification=?", "NOT_VERIFIED").Count(&countNotVerified).Error

	var countTotalNotRegistered int64

	err = s.db.WithContext(ctx).Table("students_description").Where("verification=?", "NOT_REGISTERED").Count(&countTotalNotRegistered).Error

	return map[string]int64{
		"count_verif":        countVerif,
		"count_not_registed": countTotalNotRegistered,
		"count_not_verified": countNotVerified,
	}

}

func (s *studentDescRepository) FindAll(ctx context.Context, query model.QueryStudentDesc) ([]entities.Student, error) {
	//TODO implement me

	var results []entities.Student

	var db = s.db.WithContext(ctx).Table("students_description")
	var err error
	if query.AcademicYear != "" && query.Semester != "" {
		db = db.Where("academic_year=? AND semester=?", query.AcademicYear, query.Semester)
	}

	if query.Name != "" {
		db = db.Where("first_name LIKE ? OR last_name LIKE ?", "%"+query.Name+"%", "%"+query.Name+"%")
	}

	if query.Verified != "" {
		db = db.Where("verification=?", query.Verified)
	}
	if query.VerifiedSKL != "" {
		db = db.Where("verification_skl=?", query.VerifiedSKL)
	}

	err = db.Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil

}

func (s *studentDescRepository) FindDocsByStudentId(ctx context.Context, id string) (entities.Document, error) {

	var result entities.Document

	err := s.db.WithContext(ctx).Where("id_student=?", id).Table("student_documents").First(&result).Error

	if err != nil {
		return entities.Document{}, err
	}

	return result, nil

}

func (s *studentDescRepository) UploadDocs(ctx context.Context, data entities.Document, isDocumentAlreadyCreated bool) error {

	if isDocumentAlreadyCreated {

		err := s.db.WithContext(ctx).Model(&data).Table("student_documents").Where("id_student=?", data.StudentId).Updates(&data).Error

		if err != nil {
			return err
		}

		return nil

	}

	err := s.db.WithContext(ctx).Table("student_documents").Create(&data).Error

	if err != nil {
		return err
	}

	return nil

}

func (u *studentDescRepository) Create(entity entities.Student, ctx context.Context) error {

	err := u.db.WithContext(ctx).Table("students_description").Create(&entity).Error

	if err != nil {
		return err
	}

	return nil

}

func (u *studentDescRepository) FindByEmail(ctx context.Context, email string) (entities.Student, error) {
	var result entities.Student

	err := u.db.WithContext(ctx).Table("students_description").Where("email=?", email).First(&result).Error

	if err != nil {

		return entities.Student{}, err
	}

	return result, nil

}

func (u *studentDescRepository) FindOneById(ctx context.Context, id string) (entities.Student, error) {

	var result entities.Student = entities.Student{}

	err := u.db.WithContext(ctx).Table("students_description").Where("student_id=?", id).First(&result).Error
	fmt.Print(err)
	if err != nil {

		return entities.Student{}, err
	}

	return result, nil

}

func (s *studentDescRepository) Update(ctx context.Context, student entities.Student, id string) error {

	err := s.db.WithContext(ctx).Table("students_description").Where("student_id=?", id).Updates(&student).Error

	if err != nil {
		return err
	}

	return nil

}
