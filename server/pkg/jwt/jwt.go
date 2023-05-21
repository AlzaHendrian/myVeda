package jwtToken

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = os.Getenv("SECRET_KEY")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Unexpected singing Method : %v", token.Header["alg"])
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
