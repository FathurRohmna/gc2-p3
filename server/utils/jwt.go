package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(userID string, secretKey string, expiryDuration time.Duration) (string, error) {
	claims := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiryDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates a JWT token and returns the user ID if valid
func ValidateToken(tokenString string, secretKey string) (string, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New("invalid token signature")
		}
		return "", fmt.Errorf("invalid token: %v", err)
	}
	if !token.Valid {
		return "", errors.New("token expired")
	}
	return claims.UserID, nil
}
