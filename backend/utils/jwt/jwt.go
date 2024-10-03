package jwtToken

import (
	"fmt"
	// "os"
	"time"

	"github.com/rchmachina/bpjs-tk/model"

	"github.com/golang-jwt/jwt"
)

var SecretKey = "supperrrsecret"

func GenerateToken(user *model.LoginResponse) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &model.CustomClaims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOK := token.Claims.(jwt.MapClaims)
	if isOK && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
