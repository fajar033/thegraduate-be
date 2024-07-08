package model

type QueryStudentDesc struct {
	AcademicYear string `json:"academic_year" validate:"required"`
	Semester     string `json:"semester" validate:"required"`
	Verified     string `json:"verified" validate:"required"`
	VerifiedSKL  string `json:"verifiedSKL"`
	Name         string `json:"name"`
}
