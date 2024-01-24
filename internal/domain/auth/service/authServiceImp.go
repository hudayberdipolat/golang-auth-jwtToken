package service

import (
	"errors"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/dto"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/repository"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"
	"github.com/hudayberdipolat/golang-auth-jwtToken/pkg/jwtToken"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authServiceImp struct {
	authRepo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return authServiceImp{
		authRepo: repo,
	}
}

func (a authServiceImp) Register(registerRequest dto.RegisterRequest) (*dto.AuthResponse, error) {
	//find username
	if checkUsername := a.authRepo.CheckUsername(registerRequest.Username); checkUsername != true {
		return nil, errors.New("Bu username ady eýýäm ulanylýar!!!")
	}
	// hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 16)
	// create user model type
	user := models.User{
		Username:  registerRequest.Username,
		FullName:  registerRequest.FullName,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// create user
	if err := a.authRepo.Create(user); err != nil {
		return nil, err
	}
	getUser, err := a.authRepo.GetByUserWithUsername(registerRequest.Username)
	if err != nil {
		return nil, err
	}
	// generate token
	accessToken, errToken := jwtToken.GenerateToken(getUser.ID, getUser.Username)
	if errToken != nil {
		return nil, errToken
	}
	registerResponse := dto.NewAuthResponse(*getUser, accessToken)
	// return response
	return &registerResponse, nil
}

func (a authServiceImp) Login(loginRequest dto.LoginRequest) (*dto.AuthResponse, error) {
	//get user
	getUser, err := a.authRepo.GetByUserWithUsername(loginRequest.Username)
	if err != nil {
		return nil, err
	}
	// password composer
	errPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginRequest.Password))
	if errPassword != nil {
		return nil, err
	}
	//generate token
	accessToken, errToken := jwtToken.GenerateToken(getUser.ID, getUser.Username)
	if errToken != nil {
		return nil, errToken
	}
	// login response
	loginResponse := dto.NewAuthResponse(*getUser, accessToken)
	return &loginResponse, nil
}
