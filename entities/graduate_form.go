package entities

type GraduateForm struct {
	StudentID        string  `gorm:"column:student_id" json:"student_id"`
	Id               string  `json:"id"`
	FullName         string  `gorm:"column:full_name" json:"full_name"`
	BirthDate        string  `gorm:"column:birth_date" json:"birth_date"`
	BirthPlace       string  `gorm:"column:birth_place" json:"birth_place"`
	Gender           string  `gorm:"column:gender" json:"gender"`
	Address          string  `gorm:"column:address" json:"address"`
	PhoneNumber      string  `gorm:"column:phone_number" json:"phone_number"`
	Major            string  `gorm:"column:major" json:"major"`
	GPA              float32 `gorm:"column:gpa" json:"gpa"`
	Religion         string  `gorm:"column:religion" json:"religion"`
	Level            string  `gorm:"column:level" json:"level"`
	FatherName       string  `gorm:"column:dad" json:"dad"`
	MotherName       string  `json:"mother" gorm:"column:mother"`
	ParentTelp       string  `gorm:"column:parent_telp" json:"parent_telp"`
	ParentAddress    string  `gorm:"column:parent_address" json:"parent_address"`
	CommencementDate string  `gorm:"column:commencement_date" json:"commencement_date"`
}
