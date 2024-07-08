package entities

type DocumentAdminEntity struct {
	Id                       string `json:"id,omitempty" gorm:"column:id"`
	Nidn                     string `json:"nidn,omitempty" gorm:"column:nidn"`
	StudentId                string `json:"student_id,omitempty" gorm:"column:student_id"`
	AdvisorAssignmentLetter  string `json:"advisor_assignment_letter,omitempty" gorm:"column:advisor_assignment_letter"`
	ExaminerAssignmentLetter string `json:"examiner_assignment_letter,omitempty" gorm:"column:examiner_assignment_letter"`
	Invitation               string `json:"invitation,omitempty" gorm:"column:invitation"`
	TempGrad                 string `json:"temp_grad,omitempty" gorm:"column:temp_grad"`
	OfficialReport           string `json:"official_report,omitempty"`
}

type StudentDocumentAdmin struct {
	Student  Student             `json:"student"`
	Document DocumentAdminEntity `json:"document"`
}
