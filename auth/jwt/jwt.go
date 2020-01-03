package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// JWT struct
type JWT struct {
	SigningKey []byte
}

// CustomClaims custom claims
type CustomClaims struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
