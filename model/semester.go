package model

type CreateSemesterModel struct {
	AcademicYear string `json:"academic_year"`
	Semester     string `json:"semester"`
}
