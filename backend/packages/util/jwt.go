package util

import (
	"go-todo/packages/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	// add other fields here if needed
}

func GenerateJwt(userId int) (string, error) {

	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		jwt.RegisteredClaims{
			Issuer:    "go-todo",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Subject:   strconv.Itoa(userId),
		},
	})

	token, err := tokenStruct.SignedString([]byte(config.Get().JwtSecret))

	return token, err
}

func ParseJwt(token string) (JwtClaims, error) {
	var claims JwtClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().JwtSecret), nil
	})
	return claims, err
}
