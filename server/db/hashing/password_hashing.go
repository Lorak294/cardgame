package hashing

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w",err)
	}
    return string(bytes), nil
}

// VerifyPassword verifies if the given password matches the stored hash.
func CheckPassword(password string, hash string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
