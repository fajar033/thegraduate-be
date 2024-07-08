package semester

import (
	"errors"
	"thegraduate-server/entities"
	"thegraduate-server/model"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type SemesterRepository struct {
	gorm *gorm.DB
}

func NewSemesterRepository(gorm *gorm.DB) SemesterRepository {
	return SemesterRepository{gorm}
}

func (s *SemesterRepository) FindById(id string) *entities.SemesterEntity {

	var result entities.SemesterEntity
	err := s.gorm.Table("semester").Where("id=?", id).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: err.Error()})
		}
		panic(err)
	}

	return &result

}

func (s *SemesterRepository) Delete(id string) {

	var result entities.SemesterEntity
	err := s.gorm.Table("semester").Where("id=?", id).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: err.Error()})
		}
		panic(err)
	}

	err = s.gorm.Table("semester").Where("id=?", id).Delete(&entities.SemesterEntity{}).Error
	if err != nil {
		panic(err)
	}

}

func (s *SemesterRepository) CreateSemester(payload model.CreateSemesterModel) {

	id, _ := gonanoid.New(10)

	var semester entities.SemesterEntity = entities.SemesterEntity{
		Id:           id,
		AcademicYear: payload.AcademicYear,
		Semester:     payload.Semester,
		Status:       "AKTIF",
	}

	err := s.gorm.Exec("UPDATE semester SET status = 'NON-AKTIF' WHERE id != '1'").Error
	if err != nil {
		panic(errors.New("ERROR WHILE UPDATING"))
	}

	err = s.gorm.Table("semester").Create(&semester).Error

	if err != nil {
		panic(err)
	}
}

func (s *SemesterRepository) GetSemesters() []entities.SemesterEntity {

	var result []entities.SemesterEntity
	err := s.gorm.Table("semester").Find(&result).Error

	if err != nil {
		panic(err)
	}

	return result

}
