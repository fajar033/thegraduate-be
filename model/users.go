package model

type UserModel struct {
	Id        string `json:"id" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=student admin lecturer"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type LoginModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type EmailModel struct {
	Email string `json:"email"`
}

type Forgetpassword struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Session  string `json:"session"`
}
