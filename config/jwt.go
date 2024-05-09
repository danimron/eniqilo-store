package config

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	Name    string
	StaffId int
	jwt.RegisteredClaims
}
