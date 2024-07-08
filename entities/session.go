package entities

type SessionEntity struct {
	SessionCode string `gorm:"session_code"`
	Username    string `gorm:"username"`
}
