package jwtToken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

var SecretKey = []byte("das#jd!ahDjSwr$we$ry$wbw_we^t*&^$%^#$sa)s")

func GenerateToken(userID int, username string) (string, error) {
	// Create the claims
	claims := Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // Token expires in 2 hour
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function to verify a JWT token

func VerifyToken(tokenString string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Check if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
