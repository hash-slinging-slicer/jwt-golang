package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("tyadguhsjikdahgsdha2345675ertdfgv")

type JWTClaim struct {
	Id       int64
	Nama     string
	Username string
	jwt.RegisteredClaims
}
