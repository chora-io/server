package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a jwt token using the provided secret and subject.
func GenerateJWT(secret string, sub string) (string, error) {

	// generate new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iss": "chora-web-user", // TODO
		"aud": "chora-web-user", // TODO
	})

	// sign token with secret application key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// return token string
	return tokenString, nil
}

// ValidateJWT validates a jwt token using the provided secret and token. If valid,
// the subject from the token claims is returned.
func ValidateJWT(secret string, token string) (string, error) {

	// parse token using secret application key
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	// return error if token is invalid
	if !t.Valid {
		return "", errors.New("invalid")
	}

	// return subject if token is valid
	return t.Claims.GetSubject()
}
