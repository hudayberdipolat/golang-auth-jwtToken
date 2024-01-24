package repository

import "github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"

type UserRepository interface {
	GetUser(username string) (*models.User, error)
	GetUserByID(userID int) (*models.User, error)
	ChangeUserData(userID int, user models.User) error
	ChangeUserPassword(userID int, password string) error
}
