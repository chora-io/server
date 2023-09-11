package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a jwt token using a secret application key
func GenerateJWT(secret string, address string, signature string) (string, error) {

	// generate new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"address":   address,
		"signature": signature,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	// sign token with secret application key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// return token string
	return tokenString, nil
}

// ValidateJWT validates a jwt token using a secret application key
func ValidateJWT(secret string, token string) error {

	// parse token using secret application key
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	// return error if token is invalid
	if !t.Valid {
		return errors.New("invalid")
	}

	return nil
}
