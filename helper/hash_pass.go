package helper

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		log.Fatal("failed to hash password: %v", err)
	}
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}