package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key_here") // Make sure to use a secure key

func GenerateJWT(userID string, duration time.Duration) (string, error) {
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(duration).Unix(),
        Issuer:    userID,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
    claims := &jwt.StandardClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("invalid token")
    }
    return claims, nil
}