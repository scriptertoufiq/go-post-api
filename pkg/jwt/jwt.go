package jwt

import (
	"errors"
	"fmt"
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

// ValidateToken parses and verifies an access token, returning the id and
// username claims embedded by CreateToken. Expiry is enforced by jwt.Parse.
func ValidateToken(tokenString, secretKey string) (int64, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	// numeric claims round-trip through JSON as float64
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, "", errors.New("invalid token: missing id claim")
	}
	username, _ := claims["username"].(string)

	return int64(id), username, nil
}
