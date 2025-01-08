package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	log.Printf("failed to hash password due to %s", err)
	return string(hashedPassword), true
}

func ValidateHash(hashedpassword []byte, plainTextPassword string) error {
	if err := bcrypt.CompareHashAndPassword(
		hashedpassword,
		[]byte(plainTextPassword)); err != nil {
		return Unauthorized
	}

	return nil
}
