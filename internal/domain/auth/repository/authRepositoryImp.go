package repository

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"
	"gorm.io/gorm"
)

type authRepositoryImp struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepositoryImp{
		db: db,
	}
}

func (a authRepositoryImp) Create(user models.User) error {
	if err := a.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (a authRepositoryImp) GetByUserWithUsername(username string) (*models.User, error) {
	var user models.User
	if err := a.db.Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a authRepositoryImp) CheckUsername(username string) bool {
	var user models.User
	a.db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return true
	}
	return false
}
