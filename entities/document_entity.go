package entities

type Document struct {
	Id                    string `json:"id,omitempty" gorm:"primaryKey"`
	Toeic                 string `json:"toeic_certificate,omitempty" gorm:"column:toeic_certificate"`
	IdCard                string `json:"id_card,omitempty" gorm:"column:id_card"`
	FamilyCard            string `json:"family_card,omitempty" gorm:"column:family_card"`
	ThesisFile            string `json:"thesis_file,omitempty" gorm:"column:thesis_file"`
	TempGradCertificate   string `json:"temp_graduation_certificate,omitempty" gorm:"column:temp_graduation_certificate"`
	GradCertificate       string `json:"graduation_certificate,omitempty" gorm:"column:graduation_certificate"`
	Photo                 string `json:"photo,omitempty" gorm:"column:photo"`
	CompetencyCertificate string `json:"competency_certificate,omitempty" gorm:"column:competency_certificate"`
	Article               string `json:"article,omitempty" gorm:"column:article"`
	StudentCard           string `json:"student_card,omitempty" gorm:"column:student_card"`
	ValiditySheet         string `json:"validity_sheet,omitempty" gorm:"column:validity_sheet"`
	StudentId             string `json:"id_student,omitempty" gorm:"column:id_student"`
	BirthCertificate      string `json:"birth_certificate,omitempty" gorm:"column:birth_certificate"`
}
