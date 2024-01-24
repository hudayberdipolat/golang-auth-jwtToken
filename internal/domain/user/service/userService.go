package service

import "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/dto"

type UserService interface {
	GetUser(username string) (*dto.UserResponse, error)
	UpdateUserData(userID int, updateUserData dto.ChangeUserData) error
	UpdateUserPassword(userID int, updateUserPassword dto.ChangeUserPassword) error
}
