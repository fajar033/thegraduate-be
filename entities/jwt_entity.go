package entities

import "github.com/golang-jwt/jwt/v5"

type JwtClaimsCustom struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	Npm       string `json:"npm"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}
