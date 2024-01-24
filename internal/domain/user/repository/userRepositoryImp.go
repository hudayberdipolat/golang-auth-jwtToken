package repository

import (
	"errors"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"
	"gorm.io/gorm"
)

type userRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: db,
	}
}

func (u userRepositoryImp) GetUser(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepositoryImp) ChangeUserData(userID int, user models.User) error {
	var userModel models.User
	if err := u.db.Model(&userModel).Where("id=?", userID).Updates(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu username ady eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (u userRepositoryImp) ChangeUserPassword(userID int, password string) error {
	var user models.User
	if err := u.db.Model(&user).Where("id=?", userID).Updates(models.User{
		Password: password}).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepositoryImp) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id=?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
