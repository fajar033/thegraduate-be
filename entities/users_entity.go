package entities

type UserEntity struct {
	Id        string `json:"id" gorm:"column:id"`
	Username  string `json:"username"  gorm:"column:username"`
	FirstName string `gorm:"first_name" json:"first_name"`
	LastName  string `json:"last_name" gorm:"last_name"`
	Password  string `json:"password"  gorm:"column:password"`
	Role      string `json:"role"  gorm:"column:role"`
	Email     string `json:"email"  gorm:"column:email"`
}
