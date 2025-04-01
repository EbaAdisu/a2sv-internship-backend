package Infrastructure

import (
    "time"

    "task-manager/Domain"

    "github.com/dgrijalva/jwt-go"
)

// GenerateToken creates a JWT token for a given user.
func GenerateToken(user Domain.User, expires time.Time) (string, error) {
    claims := jwt.MapClaims{
        "user_id":  user.ID,
        "username": user.Username,
        "role":     user.Role,
        "exp":      expires.Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtSecret())
}