package dto

import ("github.com/golang-jwt/jwt")


type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}