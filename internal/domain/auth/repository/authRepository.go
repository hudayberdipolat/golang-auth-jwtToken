package repository

import "github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"

type AuthRepository interface {
	Create(user models.User) error
	GetByUserWithUsername(username string) (*models.User, error)
	CheckUsername(username string) bool
}
