package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/utils/response"
	"github.com/hudayberdipolat/golang-auth-jwtToken/pkg/jwtToken"
	"net/http"
)

// Claims represents the JWT claims.

// AuthMiddleware is the JWT authentication middleware.

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		errResponse := response.Error(http.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	claims, err := verifyToken(token)
	if err != nil {
		errResponse := response.Error(http.StatusUnauthorized, "Invalid token", "Invalid token", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	ctx.Locals("username", claims.Username)
	ctx.Locals("user_id", claims.UserID)
	return ctx.Next()
}

func verifyToken(tokenString string) (*jwtToken.Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwtToken.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtToken.SecretKey, nil
	})
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Check if the token is valid
	if claims, ok := token.Claims.(*jwtToken.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
