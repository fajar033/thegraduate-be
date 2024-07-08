package entities

type Lecturer struct {
	Nidn         string `json:"nidn" gorm:"column:nidn" `
	LecturerName string `json:"lecturer_name" gorm:"column:lecturer_name"`
}
