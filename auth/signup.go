package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = ""

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Panic("Error in signing token", signedToken)
		return "", err
	}
	return signedToken, nil
}

func ParseToken(signedToken string) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Panic("Error in token parsing", err)
	}

	if !token.Valid {
		log.Panic("Invalid Token")
	}
}
