package token

import "github.com/dgrijalva/jwt-go"

type AccessCustomClaim struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
