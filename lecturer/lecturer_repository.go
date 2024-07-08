package lecturer

import (
	"context"
	"errors"
	"thegraduate-server/entities"
	"thegraduate-server/interfaces"
	"thegraduate-server/model"

	"gorm.io/gorm"
)

type lecturerRepository struct {
	db *gorm.DB
}

func NewLecturerRepository(db *gorm.DB) interfaces.ILecturerRepository {

	return &lecturerRepository{
		db: db,
	}

}

func (l *lecturerRepository) FindOneLecturerById(ctx context.Context, id string) (*entities.Lecturer, error) {

	var result *entities.Lecturer
	err :=	l.db.WithContext(ctx).Table("lecturer").Where("nidn=?", id).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&model.NotFoundError{Message: err.Error()})
		}
		panic(err)
	}

	return result, nil

}

func (l *lecturerRepository) FindAllDocuments(ctx context.Context, nidn string, name string) []entities.StudentDocumentAdmin {

	var results []entities.Student
	var err error
	db := l.db.WithContext(ctx).Table("students_description")
	
	if name != "" {
		db = db.Where("first_name LIKE ? OR last_name LIKE ?", "%"+name+"%", "%"+name+"%")

	}
	if nidn == "" {
		err = db.Find(&results).Error
	} else {
		err = db.Where("nidn_advisor_one=?", nidn).Find(&results).Error
	}

	if err != nil {
		panic(err)

	}
	var documentStudentAdmin []entities.StudentDocumentAdmin = make([]entities.StudentDocumentAdmin, len(results))
	for index, data := range results {
		var resultDocument entities.DocumentAdminEntity = entities.DocumentAdminEntity{}
		err := l.db.WithContext(ctx).Table("documents_admin").Where("student_id=?", data.StudentID).First(&resultDocument).Error

		if resultDocument.AdvisorAssignmentLetter != "" {
			resultDocument.AdvisorAssignmentLetter = "https://storage.googleapis.com/thegraduate-bucket/" + resultDocument.AdvisorAssignmentLetter
		}

		if resultDocument.ExaminerAssignmentLetter != "" {

			resultDocument.ExaminerAssignmentLetter = "https://storage.googleapis.com/thegraduate-bucket/" + resultDocument.ExaminerAssignmentLetter
		}

		if resultDocument.OfficialReport != "" {

			resultDocument.OfficialReport = "https://storage.googleapis.com/thegraduate-bucket/" + resultDocument.OfficialReport
		}

		if resultDocument.Invitation != "" {

			resultDocument.Invitation = "https://storage.googleapis.com/thegraduate-bucket/" + resultDocument.Invitation
		}

		if resultDocument.TempGrad != "" {
			resultDocument.TempGrad = "https://storage.googleapis.com/thegraduate-bucket/" + resultDocument.TempGrad
		}

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}

		}
		documentStudentAdmin[index] = entities.StudentDocumentAdmin{
			Student:  data,
			Document: resultDocument,
		}

	}

	if err != nil {
		panic(err)
	}

	return documentStudentAdmin

}

func (l *lecturerRepository) GetTotalStatic(ctx context.Context) map[string]any {
	//TODO implement me

	var count int64
	err := l.db.WithContext(ctx).Table("lecturer").Count(&count).Error

	var countAdvisorLetter int64

	err = l.db.WithContext(ctx).Table("documents_admin").Where("advisor_assignment_letter IS NOT NULL").Count(&countAdvisorLetter).Error

	var countExaminerLetter int64

	err = l.db.WithContext(ctx).Table("documents_admin").Where("examiner_assignment_letter IS NOT NULL").Count(&countExaminerLetter).Error

	var countInvitationLetter int64

	err = l.db.WithContext(ctx).Table("documents_admin").Where("invitation IS NOT NULL").Count(&countInvitationLetter).Error

	var countTempGrade int64
	err = l.db.WithContext(ctx).Table("documents_admin").Where("temp_grad IS NOT NULL").Count(&countTempGrade).Error

	var officialReport int64

	err = l.db.WithContext(ctx).Table("documents_admin").Where("official_report IS NOT NULL").Count(&officialReport).Error

	if err != nil {
		panic(err)
	}

	var data map[string]any = map[string]any{
		"total_lecturer": count,
		"total_examiner_lette": countExaminerLetter,
		"total_invitation": countInvitationLetter,
		"total_official_report": officialReport,
		"total_advisor_letter": countAdvisorLetter,
		"total_temp_grad": countTempGrade,
	}

	return data
}




func (l *lecturerRepository) FindDocumentByStudentAndNidn(ctx context.Context, studentId string, nidn string) (entities.DocumentAdminEntity, error) {
	//TODO implement me
	var result entities.DocumentAdminEntity
	err := l.db.WithContext(ctx).Table("documents_admin").Where("student_id=? AND nidn=?", studentId, nidn).First(&result).Error

	if err != nil {
		return entities.DocumentAdminEntity{}, err
	}
	return result, nil
}

func (l *lecturerRepository) FindOneById(ctx context.Context, id string) (map[string]interface{}, error) {
	//TODO implement me
	var result map[string]interface{}
	err := l.db.WithContext(ctx).Where("id = ? AND role = ?", id, "lecturer").Table("users").Find(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (l *lecturerRepository) UploadDocsAdmin(ctx context.Context, data entities.DocumentAdminEntity, isDocumentAlreadyCreated bool) {
	//TODO implement me
	if isDocumentAlreadyCreated {
		err := l.db.WithContext(ctx).Table("documents_admin").Where("student_id=?", data.StudentId).Updates(&data).Error
		if err != nil {

			panic(err)
		}

		return

	}

	err := l.db.WithContext(ctx).Table("documents_admin").Create(&data).Error

	if err != nil {
		panic(err)
	}

}

func (l *lecturerRepository) List(ctx context.Context) []map[string]interface{} {
	var result []map[string]interface{}

	err := l.db.WithContext(ctx).Table("lecturer").Find(&result).Error

	if err != nil {
		panic(err)
	}

	return result

}

func (l *lecturerRepository) Insert(ctx context.Context, data entities.Lecturer) {

	err := l.db.WithContext(ctx).Table("lecturer").Create(&data).Error

	if err != nil {
		panic(err.Error())
	}

}

func (l *lecturerRepository) Update(ctx context.Context, data entities.Lecturer, nidn string) {

	err := l.db.WithContext(ctx).Table("lecturer").Where("nidn=?", nidn).Updates(data).Error

	if err != nil {
		panic(err)
	}

}

func (l *lecturerRepository) Delete(ctx context.Context, id string) {

	var result entities.Lecturer
	err := l.db.WithContext(ctx).Where("nidn=?", id).Table("lecturer").First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(model.NotFoundError{
				Message: "no data found ",
			})
		}
		panic(err)
	}

	err = l.db.WithContext(ctx).Table("lecturer").Where("nidn=?", id).Delete(&entities.Lecturer{}).Error

	if err != nil {
		panic(err)
	}

}
