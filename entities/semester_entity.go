package entities

type SemesterEntity struct {
	Id           string `gorm:"column:id" json:"id"`
	AcademicYear string `gorm:"column:academic_year" json:"academic_year"`
	Semester     string `gorm:"column:semester" json:"semester"`
	Status       string `gorm:"column:status" json:"status"`
}
