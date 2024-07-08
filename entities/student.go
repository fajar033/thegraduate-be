package entities

type Student struct {
	StudentID        string  `gorm:"column:student_id" json:"student_id"`
	FirstName        string  `gorm:"column:first_name" json:"first_name"`
	LastName         string  `gorm:"column:last_name" json:"last_name"`
	Email            string  `gorm:"column:email" json:"email"`
	Major            string  `gorm:"column:major" json:"major"`
	NationalID       string  `gorm:"column:national_id" json:"national_id"`
	Address          string  `gorm:"column:address" json:"address"`
	MessageSKL       string  `json:"message_skl" gorm:"column:message_skl"`
	VerificationSKL  string  `json:"verification_skl" gorm:"column:verification_skl"`
	BirthDate        string  `gorm:"column:birth_date" json:"birth_date"`
	NidnAdvisorOne   string  `json:"nidn_advisor_one" gorm:"column:nidn_advisor_one"`
	NidnAdvisorTwo   string  `json:"nidn_advisor_two" gorm:"column:nidn_advisor_two"`
	NidnReligion     string  `json:"nidn_religion" gorm:"column:nidn_religion"`
	BirthPlace       string  `gorm:"column:birth_place" json:"birth_place"`
	PhoneNumber      string  `gorm:"column:phone_number" json:"phone_number"`
	TelephoneNumber  string  `gorm:"column:telephone_number" json:"telephone_number"`
	CreditCourse     int     `gorm:"column:credit_course" json:"credit_course"`
	GPA              float32 `gorm:"column:gpa" json:"gpa"`
	Message          string  `json:"message" gorm:"column:message"`
	Verification     string  `gorm:"column:verification" json:"verification"`
	AcademicYear     string  `gorm:"column:academic_year" json:"academic_year"`
	Semester         string  `gorm:"column:semester" json:"semester"`
	Examiner         string  `gorm:"column:examiner" json:"examiner"`
	Gender           string  `gorm:"column:gender" json:"gender"`
	ThesisTitle      string  `gorm:"column:thesis_title" json:"thesis_title"`
	Advisor          string  `gorm:"column:advisor" json:"advisor"`
	ReligionAdvisor  string  `gorm:"column:religion_advisor" json:"religion_advisor"`
	GraduateDate     string  `gorm:"column:graduate_date" json:"graduate_date"`
	CommencementDate string  `gorm:"column:commencement_date" json:"commencement_date"`
}
