package Infrastructure

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPassword compares a hashed password with its plain-text version.
func CheckPassword(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}