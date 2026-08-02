package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	// Implement token generation logic here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(60 * time.Minute).Unix(), // Token expiration time (e.g., 24 hours)
	})
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	// Implement refresh token generation logic here
	return "", nil
}
