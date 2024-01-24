package service

import (
	"errors"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/dto"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/repository"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (u userServiceImp) GetUser(username string) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewUserResponse(*user)
	return &userResponse, nil
}

func (u userServiceImp) UpdateUserData(userID int, updateUserData dto.ChangeUserData) error {
	updateUser := models.User{
		Username:  updateUserData.Username,
		FullName:  updateUserData.FullName,
		UpdatedAt: time.Now(),
	}
	if err := u.userRepo.ChangeUserData(userID, updateUser); err != nil {
		return err
	}
	return nil
}

func (u userServiceImp) UpdateUserPassword(userID int, updateUserPassword dto.ChangeUserPassword) error {
	getUser, errGetUser := u.userRepo.GetUserByID(userID)
	if errGetUser != nil {
		return errGetUser
	}
	errOldPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(updateUserPassword.OldPassword))
	if errOldPassword != nil {
		return errors.New("Köne password nädogry!!!")
	}
	if err := u.userRepo.ChangeUserPassword(userID, updateUserPassword.Password); err != nil {
		return err
	}
	return nil
}
